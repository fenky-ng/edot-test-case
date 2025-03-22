package warehouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	string_util "github.com/fenky-ng/edot-test-case/product/internal/utility/string"
)

func (r *RepoHttpWarehouse) GetStocks(ctx context.Context, input model.GetStocksInput) (output model.GetStocksOutput, err error) {
	baseURL := r.config.WarehouseRestServiceAddress + constant.WarehouseGetStocks

	params := url.Values{}
	if len(input.ProductIds) != 0 {
		params.Add("productIds", strings.Join(string_util.ParseUuidArrToStringArr(input.ProductIds), ","))
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrGetStocks, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrGetStocks, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var warehouseResp model.HttpGetStocksResponse
	err = json.NewDecoder(resp.Body).Decode(&warehouseResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetStocks, err)
		return output, err
	}

	if warehouseResp.Error != "" {
		err = errors.Join(in_err.ErrGetStocks, errors.New(warehouseResp.Error))
		return output, err
	}

	for _, item := range warehouseResp.Data {
		stock := model.ExtProductStock{
			ProductId:  item.ProductId,
			Warehouses: make([]model.ExtProductWarehouse, 0),
		}

		for _, warehouse := range item.Warehouses {
			stock.Warehouses = append(stock.Warehouses, model.ExtProductWarehouse{
				WarehouseId:     warehouse.WarehouseId,
				WarehouseStatus: warehouse.WarehouseStatus,
				Stock:           warehouse.Stock,
			})
		}

		output.Stocks = append(output.Stocks, stock)
	}

	return output, nil
}

package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	string_util "github.com/fenky-ng/edot-test-case/order/internal/utility/string"
)

func (r *RepoHttpProduct) GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error) {
	baseURL := r.config.ProductRestServiceAddress + constant.ProductGetProductsUri

	params := url.Values{}
	if len(input.Ids) != 0 {
		params.Add("ids", strings.Join(string_util.ParseUuidArrToStringArr(input.Ids), ","))
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrGetProducts, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrGetProducts, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var productResp model.HttpGetProductsResponse
	err = json.NewDecoder(resp.Body).Decode(&productResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetProducts, err)
		return output, err
	}

	if productResp.Error != "" {
		err = errors.Join(in_err.ErrGetProducts, errors.New(productResp.Error))
		return output, err
	}

	for _, item := range productResp.Data {
		var stockWarehouses []model.ExtProductWarehouse
		for _, warehouse := range item.Stock.Warehouses {
			stockWarehouses = append(stockWarehouses, model.ExtProductWarehouse{
				WarehouseId:     warehouse.WarehouseId,
				WarehouseStatus: warehouse.WarehouseStatus,
				Stock:           warehouse.Stock,
			})
		}

		output.Products = append(output.Products, model.ExtProduct{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Status:      item.Status,
			Shop: model.ExtShop{
				Id:     item.Shop.Id,
				Name:   item.Shop.Name,
				Status: item.Shop.Status,
			},
			Stock: model.ExtProductStock{
				Total:      item.Stock.Total,
				Warehouses: stockWarehouses,
			},
		})
	}

	return output, nil
}

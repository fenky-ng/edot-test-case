package warehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (r *RepoHttpWarehouse) DeductStocks(ctx context.Context, input model.DeductStocksInput) (output model.DeductStocksOutput, err error) {
	var deductionStockItems []model.HttpDeductStockItem
	for _, item := range input.Items {
		deductionStockItems = append(deductionStockItems, model.HttpDeductStockItem{
			ProductId:   item.ProductId,
			WarehouseId: item.WarehouseId,
			Quantity:    item.Quantity,
		})
	}
	reqBody := model.HttpDeductStocksRequest{
		UserId:  input.UserId,
		OrderNo: input.OrderNo,
		Items:   deductionStockItems,
		Release: input.Release,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		err = errors.Join(in_err.ErrDeductStocks, err)
		return output, err
	}

	req, err := http.NewRequest(http.MethodPost, r.config.WarehouseRestServiceAddress+constant.WarehouseDeductStocksUri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}
	req.Header.Set(constant.HeaderApiKey, r.config.WarehouseRestServiceApiKey)
	req.Header.Set(constant.ContentType, constant.ApplicationJSON)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrDeductStocks, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrDeductStocks, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var warehouseResp model.HttpDeductStocksResponse
	err = json.NewDecoder(resp.Body).Decode(&warehouseResp)
	if err != nil {
		err = errors.Join(in_err.ErrDeductStocks, err)
		return output, err
	}

	if warehouseResp.Error != "" {
		err = errors.Join(in_err.ErrDeductStocks, errors.New(warehouseResp.Error))
		return output, err
	}

	output.Successful = warehouseResp.Successful

	return output, nil
}

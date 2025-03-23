package product

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (r *RepoHttpProduct) GetProductById(ctx context.Context, input model.GetProductByIdInput) (output model.GetProductByIdOutput, err error) {
	url := r.config.ProductRestServiceAddress + fmt.Sprintf(constant.ProductGetProductByIdUri, input.Id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrGetProductById, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrGetProductById, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var productResp model.HttpGetProductByIdResponse
	err = json.NewDecoder(resp.Body).Decode(&productResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetProductById, err)
		return output, err
	}

	if productResp.Error != "" {
		err = errors.Join(in_err.ErrGetProductById, errors.New(productResp.Error))
		return output, err
	}

	output.Id = productResp.Id
	output.Name = productResp.Name
	output.Description = productResp.Description
	output.Price = productResp.Price
	output.Status = productResp.Status
	output.Shop = model.ExtShop{
		Id:     productResp.Shop.Id,
		Name:   productResp.Shop.Name,
		Status: productResp.Shop.Status,
	}

	return output, nil
}

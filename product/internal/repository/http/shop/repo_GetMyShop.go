package shop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	"github.com/google/uuid"
)

func (r *RepoHttpShop) GetMyShop(ctx context.Context, input model.GetMyShopInput) (output model.GetMyShopOutput, err error) {
	req, err := http.NewRequest(http.MethodGet, r.config.ShopRestServiceAddress+constant.ShopGetMyShopUri, nil)
	if err != nil {
		return
	}
	req.Header.Set(constant.HeaderAuth, constant.AuthBearer+" "+input.JWT)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrGetMyShop, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrGetMyShop, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var shopResp model.HttpGetMyShopResponse
	err = json.NewDecoder(resp.Body).Decode(&shopResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetMyShop, err)
		return output, err
	}

	if shopResp.Error != "" {
		err = errors.Join(in_err.ErrGetMyShop, errors.New(shopResp.Error))
		return output, err
	}

	if shopResp.Id == uuid.Nil {
		err = errors.Join(in_err.ErrGetMyShop, errors.New("shop id is nil"))
		return output, err
	}

	output.Id = shopResp.Id
	output.Name = shopResp.Name
	output.Status = shopResp.Status

	return output, nil
}

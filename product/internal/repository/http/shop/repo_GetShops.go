package shop

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

func (r *RepoHttpShop) GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error) {
	baseURL := r.config.ShopRestServiceAddress + constant.ShopGetShopsUri

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
		err = errors.Join(in_err.ErrGetShops, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Join(in_err.ErrGetShops, fmt.Errorf("unexpected response http code: %d", resp.StatusCode))
		return output, err
	}

	var shopResp model.HttpGetShopsResponse
	err = json.NewDecoder(resp.Body).Decode(&shopResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetShops, err)
		return output, err
	}

	if shopResp.Error != "" {
		err = errors.Join(in_err.ErrGetShops, errors.New(shopResp.Error))
		return output, err
	}

	for _, item := range shopResp.Data {
		output.Shops = append(output.Shops, model.ExtShop{
			Id:     item.Id,
			Name:   item.Name,
			Status: item.Status,
		})
	}

	return output, nil
}

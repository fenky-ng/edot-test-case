package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (r *RepoHttpUser) GetUserProfile(ctx context.Context, input model.GetUserProfileInput) (output model.GetUserProfileOutput, err error) {
	req, err := http.NewRequest(http.MethodGet, r.config.UserRestServiceAddress+constant.UserGetProfileUri, nil)
	if err != nil {
		return
	}
	req.Header.Set(constant.HeaderAuth, constant.AuthBearer+" "+input.JWT)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Join(in_err.ErrGetUserProfile, err)
		return
	}
	defer resp.Body.Close()

	var authResp model.HttpGetUserProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		err = errors.Join(in_err.ErrGetUserProfile, err)
		return output, err
	}

	output.HttpCode = resp.StatusCode
	output.Error = authResp.Error
	output.Id = authResp.Id

	return output, nil
}

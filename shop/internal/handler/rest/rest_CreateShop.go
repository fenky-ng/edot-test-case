package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/shop/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/shop/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) CreateShop(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPICreateShopRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapCreateShopInput(userId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.ShopUsecase.CreateShop(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPICreateShopResponse{
		Id: out.Id,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapCreateShopInput(
	userId uuid.UUID,
	req model.RestAPICreateShopRequest,
) (output model.CreateShopInput, err error) {
	if len(req.Name) < 3 {
		err = in_err.ErrInvalidName
		return output, err
	}

	output = model.CreateShopInput{
		UserId: userId,
		Name:   req.Name,
	}
	return output, nil
}

package rest

import (
	"errors"
	"strings"

	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/shop/internal/utility/gin/response"
	string_util "github.com/fenky-ng/edot-test-case/shop/internal/utility/string"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) GetShops(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		ids []uuid.UUID
		err error
	)

	paramIds := c.Query("ids")
	if paramIds != "" {
		strIds := strings.Split(paramIds, ",")
		ids, err = string_util.ParseStringArrToUuidArr(strIds)
		if err != nil {
			err = errors.Join(in_err.ErrInvalidShopId, err)
			gin_res.ReturnError(c, err)
			return
		}
	}

	out, err := h.usecase.ShopUsecase.GetShops(ctx, model.GetShopsInput{
		Ids: ids,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	data := make([]model.RestAPIShop, 0)
	for _, shop := range out.Shops {
		data = append(data, model.RestAPIShop{
			Id:      shop.Id,
			OwnerId: shop.OwnerId,
			Name:    shop.Name,
			Status:  shop.Status,
		})
	}

	res := model.RestAPIGetShopsResponse{
		Shops: data,
	}
	gin_res.ReturnOK(c, res)
	return
}

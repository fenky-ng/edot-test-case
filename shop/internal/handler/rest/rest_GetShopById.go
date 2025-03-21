package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/shop/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) GetShopById(c *gin.Context) {
	ctx := c.Request.Context()

	shopIdStr := c.Param("shopId")
	shopId, err := uuid.Parse(shopIdStr)
	if err != nil {
		gin_res.ReturnError(c, in_err.ErrInvalidShopId)
		return
	}

	out, err := h.usecase.ShopUsecase.GetShopById(ctx, model.GetShopByIdInput{
		Id: shopId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIGetShopByIdResponse{
		RestAPIShop: model.RestAPIShop{
			Id:      out.Id,
			OwnerId: out.OwnerId,
			Name:    out.Name,
			Status:  out.Status,
		},
	}
	gin_res.ReturnOK(c, res)
	return
}

package rest

import (
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/shop/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) GetMyShop(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.ShopUsecase.GetMyShop(ctx, model.GetMyShopInput{
		UserId: userId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIGetMyShopResponse{
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

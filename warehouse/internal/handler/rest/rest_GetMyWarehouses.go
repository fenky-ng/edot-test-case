package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) GetMyWarehouses(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.GetMyWarehouses(ctx, model.GetMyWarehousesInput{
		JWT:    jwt,
		UserId: userId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	data := make([]model.RestAPIWarehouse, 0)
	for _, warehouse := range out.Warehouses {
		data = append(data, model.RestAPIWarehouse{
			Id:     warehouse.Id,
			ShopId: warehouse.ShopId,
			Name:   warehouse.Name,
			Status: warehouse.Status,
		})
	}

	res := model.RestAPIGetMyWarehousesResponse{
		Warehouses: data,
	}
	gin_res.ReturnOK(c, res)
	return
}

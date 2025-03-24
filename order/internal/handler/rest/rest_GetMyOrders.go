package rest

import (
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/order/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) GetMyOrders(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.OrderUsecase.GetOrders(ctx, model.GetOrdersInput{
		UserId: userId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	data := make([]model.RestAPIOrder, 0)
	for _, order := range out.Orders {
		data = append(data, model.RestAPIOrder{
			Id:           order.Id,
			UserId:       order.UserId,
			OrderNo:      order.OrderNo,
			Status:       order.Status,
			PaymentRefNo: order.PaymentRefNo,
		})
	}

	res := model.RestAPIGetOrdersResponse{
		Orders: data,
	}
	gin_res.ReturnOK(c, res)
	return
}

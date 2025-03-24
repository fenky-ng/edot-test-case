package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/order/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/order/internal/utility/gin/response"
)

func (h *RestAPI) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPICreateOrderRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapCreateOrderInput(jwt, userId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.OrderUsecase.CreateOrder(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPICreateOrderResponse{
		OrderNo: out.OrderNo,
		Status:  out.Status,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapCreateOrderInput(
	jwt string,
	userId uuid.UUID,
	req model.RestAPICreateOrderRequest,
) (output model.CreateOrderInput, err error) {
	if len(req.Items) == 0 {
		err = in_err.ErrNoOrderItem
		return output, err
	}

	output = model.CreateOrderInput{
		JWT:    jwt,
		UserId: userId,
		Items:  make([]model.OrderItem, 0),
	}

	for _, item := range req.Items {
		if item.ProductId == uuid.Nil {
			err = in_err.ErrInvalidProductId
			return output, err
		}
		if item.WarehouseId == uuid.Nil {
			err = in_err.ErrInvalidWarehouseId
			return output, err
		}
		if item.Quantity < 1 {
			err = in_err.ErrInvalidOrderQuantity
			return output, err
		}
		output.Items = append(output.Items, model.OrderItem{
			ProductId:   item.ProductId,
			WarehouseId: item.WarehouseId,
			Quantity:    item.Quantity,
		})
	}

	return output, nil
}

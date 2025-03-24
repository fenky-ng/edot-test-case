package rest

import (
	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/order/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/order/internal/utility/gin/response"
	pointer_util "github.com/fenky-ng/edot-test-case/order/internal/utility/pointer"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) ConfirmPayment(c *gin.Context) {
	ctx := c.Request.Context()

	// TODO HMAC authentication

	var req model.RestAPIConfirmPaymentRequest
	err := gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	err = validateConfirmPaymentRequest(req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.OrderUsecase.UpdateOrder(ctx, model.UpdateOrderInput{
		OrderNo:      req.OrderNo,
		Status:       pointer_util.PointerOf(constant.OrderStatus_Paid),
		PaymentRefNo: pointer_util.PointerOf(req.PaymentRefNo),
		UpdatedBy:    "<payment-gateway-id>",
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIConfirmPaymentResponse{
		Successful: out.Successful,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateConfirmPaymentRequest(req model.RestAPIConfirmPaymentRequest) (err error) {
	if req.OrderNo == "" {
		return in_err.ErrMissingOrderNo
	}
	if req.PaymentRefNo == "" {
		return in_err.ErrMissingPaymentRefNo
	}
	return nil
}

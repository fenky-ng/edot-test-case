package rest

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
	managermw "github.com/fenky-ng/edot-test-case/order/internal/utility/middleware/http/httpmmanager"
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	manager := managermw.New(router)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPost,
		Path:              "/api/v1/orders",
		UseAuthentication: true,
	}, api.CreateOrder)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodGet,
		Path:              "/api/v1/orders/me",
		UseAuthentication: true,
	}, api.GetMyOrders)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodPost,
		Path:   "/api/v1/orders/payment-confirmation",
	}, api.ConfirmPayment)
}

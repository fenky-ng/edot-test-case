package rest

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	managermw "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/middleware/http/httpmmanager"
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	manager := managermw.New(router)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPost,
		Path:              "/api/v1/warehouses",
		UseAuthentication: true,
	}, api.CreateWarehouse)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPut,
		Path:              "/api/v1/warehouses/:warehouseId",
		UseAuthentication: true,
	}, api.UpdateWarehouse)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodGet,
		Path:              "/api/v1/warehouses/me",
		UseAuthentication: true,
	}, api.GetMyWarehouses)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPut,
		Path:              "/api/v1/warehouses/stocks",
		UseAuthentication: true,
	}, api.CreateOrUpdateStock)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodGet,
		Path:   "/api/v1/warehouses/stocks",
	}, api.GetStocks)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodPost,
		Path:   "/api/v1/warehouses/stocks/deduct",
	}, api.DeductStocks)
}

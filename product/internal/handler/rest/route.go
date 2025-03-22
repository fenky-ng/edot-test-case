package rest

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
	managermw "github.com/fenky-ng/edot-test-case/product/internal/utility/middleware/http/httpmmanager"
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	manager := managermw.New(router)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPost,
		Path:              "/api/v1/products",
		UseAuthentication: true,
	}, api.CreateProduct)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodGet,
		Path:              "/api/v1/products/me",
		UseAuthentication: true,
	}, api.GetMyProducts)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodGet,
		Path:   "/api/v1/products",
	}, api.GetProducts)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodGet,
		Path:   "/api/v1/products/:productId",
	}, api.GetProductById)
}

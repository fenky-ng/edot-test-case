package rest

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	managermw "github.com/fenky-ng/edot-test-case/shop/internal/utility/middleware/http/httpmmanager"
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	manager := managermw.New(router)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodPost,
		Path:              "/api/v1/shops",
		UseAuthentication: true,
	}, api.CreateShop)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodGet,
		Path:              "/api/v1/shops/me",
		UseAuthentication: true,
	}, api.GetMyShop)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodGet,
		Path:   "/api/v1/shops",
	}, api.GetShops)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodGet,
		Path:   "/api/v1/shops/:shopId",
	}, api.GetShopById)
}

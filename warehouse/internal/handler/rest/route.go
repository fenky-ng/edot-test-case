package rest

import (
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	// TODO

	// manager := managermw.New(router)

	// manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
	// 	Method:            http.MethodPost,
	// 	Path:              "/api/v1/warehouses/me",
	// 	UseAuthentication: true,
	// }, api.GetMyWarehouses)
}

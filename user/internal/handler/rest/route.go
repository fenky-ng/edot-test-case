package rest

import (
	"net/http"

	"github.com/fenky-ng/edot-test-case/user/internal/model"
	managermw "github.com/fenky-ng/edot-test-case/user/internal/utility/middleware/http/httpmmanager"
	"github.com/gin-gonic/gin"
)

func AssignRoute(
	api *RestAPI,
	router *gin.Engine,
) {
	manager := managermw.New(router)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodPost,
		Path:   "/api/v1/users/register",
	}, api.Register)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method: http.MethodPost,
		Path:   "/api/v1/users/login",
	}, api.Login)

	manager.AddEndpoint(model.HTTPMiddlewareManagerRequest{
		Method:            http.MethodGet,
		Path:              "/api/v1/users/me",
		UseAuthentication: true,
	}, api.GetProfile)
}

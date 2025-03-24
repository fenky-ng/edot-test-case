package httpmmanager

import (
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	"github.com/fenky-ng/edot-test-case/order/internal/utility/middleware/http/httpm"
	"github.com/fenky-ng/edot-test-case/order/internal/utility/middleware/http/httpmauth"
	"github.com/fenky-ng/edot-test-case/order/internal/utility/middleware/http/httpmlog"
	"github.com/gin-gonic/gin"
)

type MiddlewareManager struct {
	router      *gin.Engine
	middlewares []gin.HandlerFunc
}

func New(router *gin.Engine, middlewares ...gin.HandlerFunc) *MiddlewareManager {
	return &MiddlewareManager{
		router:      router,
		middlewares: middlewares,
	}
}

func (m *MiddlewareManager) AddEndpoint(endpoint model.HTTPMiddlewareManagerRequest, handler gin.HandlerFunc) {
	var handlers []gin.HandlerFunc

	// default middleware
	handlers = append(handlers, httpm.Middleware())
	handlers = append(handlers, httpmlog.Middleware())

	// custom middleware
	for _, mw := range m.middlewares {
		handlers = append(handlers, mw)
	}

	// on demand middleware
	if endpoint.UseAuthentication {
		handlers = append(handlers, httpmauth.Middleware())
	}

	// endpoint handler
	handlers = append(handlers, handler)

	m.router.Handle(endpoint.Method, endpoint.Path, handlers...)
}

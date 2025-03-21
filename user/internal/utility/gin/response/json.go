package response

import (
	"net/http"

	in_error "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, err error) {
	httpCode, errorMessage := in_error.GetHttpCodeAndErrorMessage(err)
	c.Error(err)
	c.JSON(httpCode, model.RestAPIErrorResponse{
		Error: errorMessage,
	})
}

func ReturnOK(c *gin.Context, res any) {
	c.JSON(http.StatusOK, res)
}

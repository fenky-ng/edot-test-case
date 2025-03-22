package httpmauth

import (
	"net/http"
	"strings"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/usecase"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constant.HeaderAuth)
		if authHeader == "" {
			gin_res.ReturnError(c, in_err.ErrMissingAuthToken)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != constant.AuthBearer {
			gin_res.ReturnError(c, in_err.ErrInvalidAuthTokenFormat)
			c.Abort()
			return
		}

		token := parts[1]
		if token == "" {
			gin_res.ReturnError(c, in_err.ErrMissingAuthToken)
			c.Abort()
			return
		}

		authOut, err := usecase.Auth.GetUserProfile(c.Request.Context(), model.GetUserProfileInput{
			JWT: token,
		})
		if err != nil {
			gin_res.ReturnError(c, err)
			c.Abort()
			return
		}
		if authOut.HttpCode != http.StatusOK || authOut.Error != "" {
			c.JSON(authOut.HttpCode, model.RestAPIErrorResponse{
				Error: authOut.Error,
			})
			c.Abort()
			return
		}

		c.Set(constant.JwtKey, token)
		c.Set(constant.UserIdKey, authOut.Id)

		// process request
		c.Next()
	}
}

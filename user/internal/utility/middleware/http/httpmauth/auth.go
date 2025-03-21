package httpmauth

import (
	"strings"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	gin_res "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/response"
	jwt_util "github.com/fenky-ng/edot-test-case/user/internal/utility/jwt"
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

		jwt, err := jwt_util.VerifyJWT(token)
		if err != nil {
			gin_res.ReturnError(c, err)
			c.Abort()
			return
		}

		c.Set(constant.UserIdKey, jwt.Sub)

		// process request
		c.Next()
	}
}

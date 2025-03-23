package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) validateApiKey(c *gin.Context) error {
	apiKey := c.GetHeader(constant.HeaderApiKey)
	if apiKey == "" {
		return in_err.ErrMissingApiKey
	}
	if apiKey != h.config.ApiKeyExternalOrder {
		return in_err.ErrInvalidApiKey
	}

	return nil
}

func getJwt(c *gin.Context) (jwt string, err error) {
	val, exists := c.Get(constant.JwtKey)
	if !exists {
		err = in_err.ErrNoJWT
		return jwt, err
	}

	jwt = val.(string)

	return jwt, nil
}

func getUserId(c *gin.Context) (userId uuid.UUID, err error) {
	val, exists := c.Get(constant.UserIdKey)
	if !exists {
		err = in_err.ErrNoUserId
		return userId, err
	}

	userId, ok := val.(uuid.UUID)
	if !ok {
		err = in_err.ErrInvalidUserId
		return userId, err
	}

	return userId, nil
}

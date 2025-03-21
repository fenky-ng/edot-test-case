package request

import (
	"errors"
	"io"

	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/gin-gonic/gin"
)

func BindRequestBodyJSON(c *gin.Context, target any) error {
	if err := c.ShouldBindBodyWithJSON(&target); err != nil {
		appErr := in_err.ErrInvalidRequestBody
		if errors.Is(err, io.EOF) {
			appErr = in_err.ErrMissingRequestBody
		}
		err = errors.Join(appErr, err)
		return err
	}
	return nil
}

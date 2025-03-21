package rest

import (
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) GetProfile(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.ProfileUsecase.GetProfile(ctx, model.GetProfileInput{
		Id: userId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIGetProfileResponse{
		Id:    out.Id,
		Name:  out.Name,
		Phone: out.Phone,
		Email: out.Email,
	}
	gin_res.ReturnOK(c, res)
	return
}

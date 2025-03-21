package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/response"
	regexp_util "github.com/fenky-ng/edot-test-case/user/internal/utility/regexp"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req model.RestAPILoginRequest
	err := gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapLoginInput(req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.AuthUsecase.Login(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPILoginResponse{
		JWT: out.JWT,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapLoginInput(req model.RestAPILoginRequest) (output model.LoginInput, err error) {
	if len(req.PhoneOrEmail) == 0 {
		err = in_err.ErrInvalidLoginRequest
		return output, err
	}

	var phone, email string
	if regexp_util.IsPhone(req.PhoneOrEmail) {
		phone = req.PhoneOrEmail
	} else if regexp_util.IsEmail(req.PhoneOrEmail) {
		email = req.PhoneOrEmail
	} else {
		err = in_err.ErrInvalidLoginRequest
		return output, err
	}

	if len(req.Password) < 6 {
		err = in_err.ErrInvalidPassword
		return output, err
	}

	output = model.LoginInput{
		Phone:    phone,
		Email:    email,
		Password: req.Password,
	}
	return output, nil
}

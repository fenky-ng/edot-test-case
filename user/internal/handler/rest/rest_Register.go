package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/user/internal/utility/gin/response"
	regexp_util "github.com/fenky-ng/edot-test-case/user/internal/utility/regexp"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var req model.RestAPIRegisterRequest
	err := gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapRegisterInput(req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.AuthUsecase.Register(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIRegisterResponse{
		Id: out.Id,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapRegisterInput(req model.RestAPIRegisterRequest) (output model.RegisterInput, err error) {
	if len(req.Name) < 3 {
		err = in_err.ErrInvalidName
		return output, err
	}

	if len(req.PhoneOrEmail) == 0 {
		err = in_err.ErrInvalidRegistrationRequest
		return output, err
	}

	var phone, email string
	if regexp_util.IsPhone(req.PhoneOrEmail) {
		phone = req.PhoneOrEmail
	} else if regexp_util.IsEmail(req.PhoneOrEmail) {
		email = req.PhoneOrEmail
	} else {
		err = in_err.ErrInvalidRegistrationRequest
		return output, err
	}

	if len(req.Password) < 6 {
		err = in_err.ErrInvalidPassword
		return output, err
	}

	output = model.RegisterInput{
		Name:     req.Name,
		Phone:    phone,
		Email:    email,
		Password: req.Password,
	}
	return output, nil
}

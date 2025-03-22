package rest

import (
	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/product/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/product/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPICreateProductRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapCreateProductInput(jwt, userId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.ProductUsecase.CreateProduct(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPICreateProductResponse{
		Id: out.Id,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapCreateProductInput(
	jwt string,
	userId uuid.UUID,
	req model.RestAPICreateProductRequest,
) (output model.CreateProductInput, err error) {
	if len(req.Name) < 3 {
		err = in_err.ErrInvalidName
		return output, err
	}

	if req.Price < 1 {
		err = in_err.ErrInvalidPrice
		return output, err
	}

	if req.Status != constant.ProductStatus_Active && req.Status != constant.ProductStatus_Inactive {
		err = in_err.ErrInvalidStatus
		return output, err
	}

	output = model.CreateProductInput{
		JWT:         jwt,
		UserId:      userId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Status:      req.Status,
	}
	return output, nil
}

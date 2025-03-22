package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) CreateOrUpdateStock(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPICreateOrUpdateStockRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapCreateOrUpdateStockInput(jwt, userId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.CreateOrUpdateStock(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPICreateOrUpdateStockResponse{
		Successful: out.Successful,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapCreateOrUpdateStockInput(
	jwt string,
	userId uuid.UUID,
	req model.RestAPICreateOrUpdateStockRequest,
) (output model.CreateOrUpdateStockInput, err error) {
	output = model.CreateOrUpdateStockInput{
		JWT:           jwt,
		UserId:        userId,
		WarehouseId:   req.WarehouseId,
		ProductId:     req.ProductId,
		Stock:         req.Stock,
		ToWarehouseId: req.ToWarehouseId,
	}
	return output, nil
}

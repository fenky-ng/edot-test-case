package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) CreateWarehouse(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPICreateWarehouseRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapCreateWarehouseInput(jwt, userId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.CreateWarehouse(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPICreateWarehouseResponse{
		Id: out.Id,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapCreateWarehouseInput(
	jwt string,
	userId uuid.UUID,
	req model.RestAPICreateWarehouseRequest,
) (output model.CreateWarehouseInput, err error) {
	if len(req.Name) < 3 {
		err = in_err.ErrInvalidName
		return output, err
	}

	if req.Status != constant.WarehouseStatus_Active && req.Status != constant.WarehouseStatus_Inactive {
		err = in_err.ErrInvalidStatus
		return output, err
	}

	output = model.CreateWarehouseInput{
		JWT:    jwt,
		UserId: userId,
		Name:   req.Name,
		Status: req.Status,
	}
	return output, nil
}

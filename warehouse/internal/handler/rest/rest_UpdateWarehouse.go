package rest

import (
	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	pointer_util "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) UpdateWarehouse(c *gin.Context) {
	ctx := c.Request.Context()

	jwt, _ := getJwt(c)

	userId, err := getUserId(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	warehouseIdStr := c.Param("warehouseId")
	warehouseId, err := uuid.Parse(warehouseIdStr)
	if err != nil {
		gin_res.ReturnError(c, in_err.ErrInvalidWarehouseId)
		return
	}

	var req model.RestAPIUpdateWarehouseRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapUpdateWarehouseInput(jwt, userId, warehouseId, req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.UpdateWarehouse(ctx, in)
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

func validateAndMapUpdateWarehouseInput(
	jwt string,
	userId uuid.UUID,
	warehouseId uuid.UUID,
	req model.RestAPIUpdateWarehouseRequest,
) (output model.UpdateWarehouseInput, err error) {
	if req.Name == nil && req.Status == nil {
		err = in_err.ErrInvalidUpdateRequest
		return output, err
	}

	if req.Name != nil && len(pointer_util.ValueOf(req.Name)) < 3 {
		err = in_err.ErrInvalidName
		return output, err
	}

	if req.Status != nil && pointer_util.ValueOf(req.Status) != constant.WarehouseStatus_Active && pointer_util.ValueOf(req.Status) != constant.WarehouseStatus_Inactive {
		err = in_err.ErrInvalidStatus
		return output, err
	}

	output = model.UpdateWarehouseInput{
		JWT:         jwt,
		UserId:      userId,
		WarehouseId: warehouseId,
		Name:        req.Name,
		Status:      req.Status,
	}
	return output, nil
}

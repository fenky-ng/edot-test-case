package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_req "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/request"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) DeductStocks(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.validateApiKey(c)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	var req model.RestAPIDeductStocksRequest
	err = gin_req.BindRequestBodyJSON(c, &req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	in, err := validateAndMapDeductStocksInput(req)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.DeductStocks(ctx, in)
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIDeductStocksResponse{
		Successful: out.Successful,
	}
	gin_res.ReturnOK(c, res)
	return
}

func validateAndMapDeductStocksInput(req model.RestAPIDeductStocksRequest) (output model.DeductStocksInput, err error) {
	output = model.DeductStocksInput{
		UserId:  req.UserId,
		OrderNo: req.OrderNo,
		Items:   make([]model.DeductStockItem, 0),
		Release: req.Release,
	}

	for _, item := range req.Items {
		if item.ProductId == uuid.Nil {
			err = in_err.ErrInvalidProductId
			return output, err
		}
		if item.WarehouseId == uuid.Nil {
			err = in_err.ErrInvalidWarehouseId
			return output, err
		}
		if item.Quantity < 1 {
			err = in_err.ErrInvalidQuantityDeduction
			return output, err
		}
		output.Items = append(output.Items, model.DeductStockItem{
			ProductId:   item.ProductId,
			WarehouseId: item.WarehouseId,
			Quantity:    item.Quantity,
		})
	}

	return output, nil
}

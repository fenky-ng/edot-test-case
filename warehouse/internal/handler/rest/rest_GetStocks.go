package rest

import (
	"errors"
	"strings"

	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/gin/response"
	string_util "github.com/fenky-ng/edot-test-case/warehouse/internal/utility/string"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) GetStocks(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		productIds []uuid.UUID
		err        error
	)

	paramProductIds := c.Query("productIds")
	if paramProductIds != "" {
		strIds := strings.Split(paramProductIds, ",")
		productIds, err = string_util.ParseStringArrToUuidArr(strIds)
		if err != nil {
			err = errors.Join(in_err.ErrInvalidProductId, err)
			gin_res.ReturnError(c, err)
			return
		}
	}
	if len(paramProductIds) == 0 {
		err = errors.Join(in_err.ErrMinOneProductIdQueryParam, err)
		gin_res.ReturnError(c, err)
		return
	}

	out, err := h.usecase.WarehouseUsecase.GetProductStocks(ctx, model.GetProductStocksInput{
		ProductIds: productIds,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	data := make([]model.RestAPIProductStock, 0)
	for _, product := range out.ProductStocks {
		cur := model.RestAPIProductStock{
			ProductId:  product.ProductId,
			Warehouses: make([]model.RestAPIProductWarehouse, 0),
		}

		for _, warehouse := range product.Warehouses {
			cur.Warehouses = append(cur.Warehouses, model.RestAPIProductWarehouse{
				WarehouseId:     warehouse.WarehouseId,
				Stock:           warehouse.Stock,
				WarehouseStatus: warehouse.WarehouseStatus,
			})
		}

		data = append(data, cur)
	}

	res := model.RestAPIGetStocksResponse{
		Products: data,
	}
	gin_res.ReturnOK(c, res)
	return
}

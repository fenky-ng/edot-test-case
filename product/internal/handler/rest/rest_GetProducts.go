package rest

import (
	"errors"
	"strings"

	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/product/internal/utility/gin/response"
	string_util "github.com/fenky-ng/edot-test-case/product/internal/utility/string"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) GetProducts(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		ids []uuid.UUID
		err error
	)

	paramIds := c.Query("ids")
	if paramIds != "" {
		strIds := strings.Split(paramIds, ",")
		ids, err = string_util.ParseStringArrToUuidArr(strIds)
		if err != nil {
			err = errors.Join(in_err.ErrInvalidProductId, err)
			gin_res.ReturnError(c, err)
			return
		}
	}

	out, err := h.usecase.ProductUsecase.GetProducts(ctx, model.GetProductsInput{
		Ids: ids,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	data := make([]model.RestAPIProduct, 0)
	for _, product := range out.Products {
		data = append(data, model.RestAPIProduct{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Status:      product.Status,
			Shop: model.RestAPIShop{
				Id:     product.Shop.Id,
				Name:   product.Shop.Name,
				Status: product.Shop.Status,
			},
			Stock: model.RestAPIProductStock{
				Total:      product.Stock.Total,
				Warehouses: mapProductStockWarehouses(product.Stock.Warehouses),
			},
		})
	}

	res := model.RestAPIGetProductsResponse{
		Products: data,
	}
	gin_res.ReturnOK(c, res)
	return
}

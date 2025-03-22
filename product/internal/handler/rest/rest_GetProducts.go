package rest

import (
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/product/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
)

func (h *RestAPI) GetProducts(c *gin.Context) {
	ctx := c.Request.Context()

	out, err := h.usecase.ProductUsecase.GetProducts(ctx, model.GetProductsInput{})
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
			Stock: model.RestAPIStock{
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

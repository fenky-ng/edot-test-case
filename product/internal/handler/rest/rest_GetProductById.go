package rest

import (
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	gin_res "github.com/fenky-ng/edot-test-case/product/internal/utility/gin/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *RestAPI) GetProductById(c *gin.Context) {
	ctx := c.Request.Context()

	productIdStr := c.Param("productId")
	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		gin_res.ReturnError(c, in_err.ErrInvalidProductId)
		return
	}

	out, err := h.usecase.ProductUsecase.GetProductById(ctx, model.GetProductByIdInput{
		Id: productId,
	})
	if err != nil {
		gin_res.ReturnError(c, err)
		return
	}

	res := model.RestAPIGetProductByIdResponse{
		RestAPIProduct: model.RestAPIProduct{
			Id:          out.Product.Id,
			Name:        out.Product.Name,
			Description: out.Product.Description,
			Price:       out.Product.Price,
			Status:      out.Product.Status,
			Shop: model.RestAPIShop{
				Id:     out.Product.Shop.Id,
				Name:   out.Product.Shop.Name,
				Status: out.Product.Shop.Status,
			},
			Stock: model.RestAPIProductStock{
				Total:      out.Product.Stock.Total,
				Warehouses: mapProductStockWarehouses(out.Product.Stock.Warehouses),
			},
		},
	}
	gin_res.ReturnOK(c, res)
	return
}

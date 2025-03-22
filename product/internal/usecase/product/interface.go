package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

type ProductUsecaseInterface interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (output model.CreateProductOutput, err error)
	GetMyProducts(ctx context.Context, input model.GetMyProductsInput) (output model.GetMyProductsOutput, err error)
	GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error)
	GetProductById(ctx context.Context, input model.GetProductByIdInput) (output model.GetProductByIdOutput, err error)
}

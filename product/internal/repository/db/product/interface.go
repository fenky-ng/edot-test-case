package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

type RepoDbProductInterface interface {
	GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error)
	GetProduct(ctx context.Context, input model.GetProductInput) (output model.GetProductOutput, err error)
	InsertProduct(ctx context.Context, input model.InsertProductInput) (output model.InsertProductOutput, err error)
}

package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

type RepoHttpProductInterface interface {
	GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error)
}

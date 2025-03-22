package product

import (
	"context"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

type RepoHttpProductInterface interface {
	GetProductById(ctx context.Context, input model.GetProductByIdInput) (output model.GetProductByIdOutput, err error)
}

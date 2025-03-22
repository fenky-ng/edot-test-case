package warehouse

import (
	"context"

	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

type RepoHttpWarehouseInterface interface {
	GetProductWarehouses(ctx context.Context, input model.GetProductWarehousesInput) (output model.GetProductWarehousesOutput, err error)
}

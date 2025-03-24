package order

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbOrder) InsertOrderDetails(ctx context.Context, input model.InsertOrderDetailsInput) (output model.InsertOrderDetailsOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableOrderDetail)

	for _, item := range input.Items {
		stmt.NewRow().
			Set("id", uuid.New()).
			Set("order_id", input.OrderId).
			Set("product_id", item.ProductId).
			Set("warehouse_id", item.WarehouseId).
			Set("price", item.Price).
			Set("quantity", item.Quantity).
			Set("created_at", input.CreatedAt).
			Set("created_by", input.CreatedBy)
	}

	_, err = stmt.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		err = errors.Join(in_err.ErrInsertOrderDetails, err)
		return output, err
	}

	return output, nil
}

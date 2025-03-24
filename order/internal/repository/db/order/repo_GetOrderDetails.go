package order

import (
	"context"
	sql "database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (r *RepoDbOrder) GetOrderDetails(ctx context.Context, input model.GetOrderDetailsInput) (output model.GetOrderDetailsOutput, err error) {
	var item model.OrderItem

	stmt := r.sql.From(constant.TableOrderDetail+" od ").
		Join(constant.TableOrder+" o ", " o.id = od.order_id ").
		Select("product_id").To(&item.ProductId).
		Select("warehouse_id").To(&item.WarehouseId).
		Select("quantity").To(&item.Quantity).
		Where("o.order_no = ?", input.OrderNo)

	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Items = append(output.Items, item)
		item = model.OrderItem{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetOrderDetails, err)
		return output, err
	}

	return output, err
}

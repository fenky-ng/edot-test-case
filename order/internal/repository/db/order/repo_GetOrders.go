package order

import (
	"context"
	sql "database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbOrder) GetOrders(ctx context.Context, input model.GetOrdersInput) (output model.GetOrdersOutput, err error) {
	var item model.Order

	stmt := r.sql.From(constant.TableOrder).
		Select("id").To(&item.Id).
		Select("user_id").To(&item.UserId).
		Select("order_no").To(&item.OrderNo).
		Select("status").To(&item.Status).
		Select("COALESCE(payment_ref_no, '') AS payment_ref_no").To(&item.PaymentRefNo).
		Where("deleted_at IS NULL").
		OrderBy("created_at DESC")

	if input.UserId != uuid.Nil {
		stmt.Where("user_id = ?", input.UserId)
	}

	if input.OrderNo != "" {
		stmt.Where("order_no = ?", input.OrderNo)
	}

	if input.GetExpiredOrders {
		stmt.Where("status = ?", constant.OrderStatus_WaitingForPayment).
			Where("created_at < (EXTRACT(EPOCH FROM NOW())*1000) - ?", constant.OrderPaymentExpiryDuration.Milliseconds())
	}

	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Orders = append(output.Orders, item)
		item = model.Order{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetOrders, err)
		return output, err
	}

	return output, nil
}

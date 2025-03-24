package order

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (r *RepoDbOrder) UpdateOrder(ctx context.Context, input model.UpdateOrderInput) (output model.UpdateOrderOutput, err error) {
	stmt := r.sql.Update(constant.TableOrder).
		Set("updated_at", input.UpdatedAt).
		Set("updated_by", input.UpdatedBy).
		Where("order_no = ?", input.OrderNo)

	if input.Status != nil {
		stmt.Set("status", input.Status)
	}

	if input.PaymentRefNo != nil {
		stmt.Set("payment_ref_no", input.PaymentRefNo)
	}

	if input.ErrorMessage != nil {
		stmt.Set("error_message", input.ErrorMessage)
	}

	_, err = stmt.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		err = errors.Join(in_err.ErrUpdateOrder, err)
		return output, err
	}

	output.Successful = true

	return output, nil
}

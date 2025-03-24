package order

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/order/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/order/internal/error"
	"github.com/fenky-ng/edot-test-case/order/internal/model"
)

func (r *RepoDbOrder) InsertOrder(ctx context.Context, input model.InsertOrderInput) (output model.InsertOrderOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableOrder).
		Set("id", input.Id).
		Set("user_id", input.UserId).
		Set("order_no", input.OrderNo).
		Set("status", input.Status).
		Set("created_at", input.CreatedAt).
		Set("created_by", input.CreatedBy)

	_, err = stmt.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		err = errors.Join(in_err.ErrInsertOrder, err)
		return output, err
	}

	return output, nil
}

package warehouse

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (r *RepoDbWarehouse) InsertWarehouse(ctx context.Context, input model.InsertWarehouseInput) (output model.InsertWarehouseOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableWarehouse).
		Set("id", input.Id).
		Set("shop_id", input.ShopId).
		Set("name", input.Name).
		Set("status", input.Status).
		Set("created_at", input.CreatedAt).
		Set("created_by", input.CreatedBy)

	_, err = stmt.ExecAndClose(ctx, r.db)
	if err != nil {
		err = errors.Join(in_err.ErrInsertWarehouse, err)
		return output, err
	}

	output.Id = input.Id

	return output, nil
}

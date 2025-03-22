package warehouse

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
)

func (r *RepoDbWarehouse) UpdateWarehouse(ctx context.Context, input model.UpdateWarehouseInput) (output model.UpdateWarehouseOutput, err error) {
	stmt := r.sql.Update(constant.TableWarehouse).
		Set("updated_at", input.UpdatedAt).
		Set("updated_by", input.UserId).
		Where("id = ?", input.WarehouseId)

	if input.Name != nil {
		stmt.Set("name", input.Name)
	}

	if input.Status != nil {
		stmt.Set("status", input.Status)
	}

	_, err = stmt.ExecAndClose(ctx, r.db)
	if err != nil {
		err = errors.Join(in_err.ErrUpdateWarehouse, err)
		return output, err
	}

	output.Id = input.WarehouseId

	return output, nil
}

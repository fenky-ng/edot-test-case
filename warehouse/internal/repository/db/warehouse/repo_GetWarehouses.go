package warehouse

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbWarehouse) GetWarehouses(ctx context.Context, input model.GetWarehousesInput) (output model.GetWarehousesOutput, err error) {
	var item model.Warehouse

	stmt := r.sql.From(constant.TableWarehouse).
		Select("id").To(&item.Id).
		Select("shop_id").To(&item.ShopId).
		Select("name").To(&item.Name).
		Select("status").To(&item.Status).
		Where("deleted_at IS NULL")

	if input.Id != uuid.Nil {
		stmt.Where("id = ?", input.Id)
	}

	if input.ShopId != uuid.Nil {
		stmt.Where("shop_id = ?", input.ShopId)
	}

	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Warehouses = append(output.Warehouses, item)
		item = model.Warehouse{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetWarehouses, err)
		return output, err
	}

	return output, nil
}

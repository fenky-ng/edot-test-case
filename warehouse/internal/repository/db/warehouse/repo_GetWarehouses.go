package warehouse

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *RepoDbWarehouse) GetWarehouses(ctx context.Context, input model.GetWarehousesInput) (output model.GetWarehousesOutput, err error) {
	var item model.Warehouse

	stmt := r.sql.From(constant.TableWarehouse).
		Select("id").To(&item.Id).
		Select("shop_id").To(&item.ShopId).
		Select("name").To(&item.Name).
		Select("status").To(&item.Status).
		Where("deleted_at IS NULL")

	if len(input.Ids) != 0 {
		stmt.Where("id = ANY(?)", pq.Array(input.Ids))
	}

	if input.ShopId != uuid.Nil {
		stmt.Where("shop_id = ?", input.ShopId)
	}

	output.WarehouseById = make(map[uuid.UUID]model.Warehouse)
	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Warehouses = append(output.Warehouses, item)
		output.WarehouseById[item.Id] = item
		item = model.Warehouse{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetWarehouses, err)
		return output, err
	}

	return output, nil
}

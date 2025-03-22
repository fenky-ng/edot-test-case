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

func (r *RepoDbWarehouse) GetStocks(ctx context.Context, input model.GetStocksInput) (output model.GetStocksOutput, err error) {
	var item model.Stock

	stmt := r.sql.From(constant.TableStock+" s ").
		Join(constant.TableWarehouse+" w ", " w.id = s.warehouse_id AND w.deleted_at IS NULL ").
		Select("s.id").To(&item.Id).
		Select("s.product_id").To(&item.ProductId).
		Select("s.warehouse_id").To(&item.WarehouseId).
		Select("w.status").To(&item.WarehouseStatus).
		Select("s.stock").To(&item.Stock).
		Where("s.deleted_at IS NULL").
		Where("s.product_id = ANY(?)", pq.Array(input.ProductIds)).
		OrderBy("s.product_id ASC")

	output.StocksByProductId = make(map[uuid.UUID][]model.Stock)
	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Stocks = append(output.Stocks, item)
		output.StocksByProductId[item.ProductId] = append(output.StocksByProductId[item.ProductId], item)
		item = model.Stock{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetStocks, err)
		return output, err
	}

	return output, nil
}

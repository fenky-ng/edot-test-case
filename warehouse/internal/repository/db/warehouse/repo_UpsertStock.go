package warehouse

import (
	context "context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	model "github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbWarehouse) UpsertStock(ctx context.Context, input model.UpsertStockInput) (output model.UpsertStockOutput, err error) {
	// LOCK
	var stockId uuid.UUID

	err = r.sql.From(constant.TableStock).
		Select("id").To(&stockId).
		Where("warehouse_id = ?", input.WarehouseId).
		Where("product_id = ? FOR UPDATE", input.ProductId).
		QueryRowAndClose(ctx, r.UseTx(ctx))

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = errors.Join(in_err.ErrUpsertStock, err)
		return output, err
	}

	// UPSERT
	onConflictStockOperator := `EXCLUDED.stock`
	if input.IsTransfer {
		onConflictStockOperator = `stock.stock + EXCLUDED.stock`
	}

	stmt := r.sql.InsertInto(constant.TableStock).
		Set("id", input.Id).
		Set("warehouse_id", input.WarehouseId).
		Set("product_id", input.ProductId).
		Set("stock", input.Stock).
		Set("created_at", input.UpsertedAt).
		Set("created_by", input.UpsertedBy).
		Clause(fmt.Sprintf(`
			ON CONFLICT ("warehouse_id", "product_id")
			DO UPDATE SET
				"stock" = %s,
				"updated_at" = EXCLUDED.created_at,
				"updated_by" = EXCLUDED.created_by
		`, onConflictStockOperator)).
		Returning("id").To(&output.Id)

	err = stmt.QueryRowAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		err = errors.Join(in_err.ErrUpsertStock, err)
		return output, err
	}

	return output, err
}

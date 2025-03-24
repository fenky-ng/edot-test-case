package warehouse

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/fenky-ng/edot-test-case/warehouse/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/warehouse/internal/error"
	"github.com/fenky-ng/edot-test-case/warehouse/internal/model"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *RepoDbWarehouse) DeductStock(ctx context.Context, input model.DeductStockInput) (output model.DeductStockOutput, err error) {
	// LOCK
	mode := ""
	if input.NoWait {
		mode = "NOWAIT"
	}

	fu := ""
	if input.Release {
		fu = fmt.Sprintf("FOR UPDATE %s", mode)
	}

	var stockId uuid.UUID

	lockStmt := r.sql.From(constant.TableStock).
		Select("id").To(&stockId).
		Where("warehouse_id = ?", input.WarehouseId).
		Where(fmt.Sprintf("product_id = ? %s", fu), input.ProductId)
	if !input.Release {
		lockStmt.Where(fmt.Sprintf("stock >= ? FOR UPDATE %s", mode), input.Quantity)
	}

	err = lockStmt.QueryRowAndClose(ctx, r.UseTx(ctx))

	if err != nil {
		appErr := in_err.ErrDeductStock
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "55P03" {
			appErr = in_err.ErrStockLocked
		} else if errors.Is(err, sql.ErrNoRows) {
			appErr = in_err.ErrInsufficientStock
		}
		err = errors.Join(appErr, err)
		return output, err
	}

	// UPDATE
	stockOperator := "-"
	if input.Release {
		stockOperator = "+"
	}

	updateStmt := r.sql.Update(constant.TableStock).
		SetExpr("stock", fmt.Sprintf("stock %s ?", stockOperator), input.Quantity).
		Set("updated_at", input.RequestedAt).
		Set("updated_by", input.RequestedBy).
		Where("warehouse_id = ?", input.WarehouseId).
		Where("product_id = ?", input.ProductId)
	if !input.Release {
		updateStmt.Where("stock >= ?", input.Quantity)
	}

	_, err = updateStmt.ExecAndClose(ctx, r.UseTx(ctx))
	if err != nil {
		err = errors.Join(in_err.ErrDeductStock, err)
		return output, err
	}

	return output, nil
}

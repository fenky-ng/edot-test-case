package shop

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbShop) GetShop(ctx context.Context, input model.GetShopInput) (output model.GetShopOutput, err error) {
	stmt := r.sql.From(constant.TableShop).
		Select("id").To(&output.Id).
		Select("owner_id").To(&output.OwnerId).
		Select("name").To(&output.Name).
		Select("status").To(&output.Status).
		Select("COALESCE(deleted_at, 0) AS deleted_at").To(&output.DeletedAt)

	if input.ShopId != uuid.Nil {
		stmt.Where("id = ?", input.ShopId)
	}

	if input.OwnerId != uuid.Nil {
		stmt.Where("owner_id = ?", input.OwnerId)
	}

	if input.ExcludeDeleted {
		stmt.Where("deleted_at IS NULL")
	}

	err = stmt.QueryRowAndClose(ctx, r.db)
	if err != nil {
		appErr := in_err.ErrGetShop
		if errors.Is(err, sql.ErrNoRows) {
			appErr = in_err.ErrShopNotFound
		}
		err = errors.Join(appErr, err)
		return output, err
	}

	return output, nil
}

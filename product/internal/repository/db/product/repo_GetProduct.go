package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbProduct) GetProduct(ctx context.Context, input model.GetProductInput) (output model.GetProductOutput, err error) {
	stmt := r.sql.From(constant.TableProduct).
		Select("id").To(&output.Product.Id).
		Select("shop_id").To(&output.Product.Shop.Id).
		Select("name").To(&output.Product.Name).
		Select("description").To(&output.Product.Description).
		Select("price").To(&output.Product.Price).
		Select("status").To(&output.Product.Status).
		Where("deleted_at IS NULL")

	if input.Id != uuid.Nil {
		stmt.Where("id = ?", input.Id)
	}

	err = stmt.QueryRowAndClose(ctx, r.db)
	if err != nil {
		appErr := in_err.ErrGetProduct
		if errors.Is(err, sql.ErrNoRows) {
			appErr = in_err.ErrProductNotFound
		}
		err = errors.Join(appErr, err)
		return output, err
	}

	return output, nil
}

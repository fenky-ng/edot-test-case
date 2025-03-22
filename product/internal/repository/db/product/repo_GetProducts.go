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

func (r *RepoDbProduct) GetProducts(ctx context.Context, input model.GetProductsInput) (output model.GetProductsOutput, err error) {
	var item model.Product

	stmt := r.sql.From(constant.TableProduct).
		Select("id").To(&item.Id).
		Select("shop_id").To(&item.Shop.Id).
		Select("name").To(&item.Name).
		Select("description").To(&item.Description).
		Select("price").To(&item.Price).
		Select("status").To(&item.Status).
		Where("deleted_at IS NULL").
		OrderBy("name ASC")

	if input.ShopId != uuid.Nil {
		stmt.Where("shop_id = ?", input.ShopId)
	}

	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Products = append(output.Products, item)
		item = model.Product{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetProducts, err)
		return output, err
	}

	return output, nil
}

package product

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/product/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/product/internal/error"
	"github.com/fenky-ng/edot-test-case/product/internal/model"
)

func (r *RepoDbProduct) InsertProduct(ctx context.Context, input model.InsertProductInput) (output model.InsertProductOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableProduct).
		Set("id", input.Id).
		Set("shop_id", input.ShopId).
		Set("name", input.Name).
		Set("description", input.Description).
		Set("price", input.Price).
		Set("status", input.Status).
		Set("created_at", input.CreatedAt).
		Set("created_by", input.CreatedBy)

	_, err = stmt.ExecAndClose(ctx, r.db)
	if err != nil {
		err = errors.Join(in_err.ErrInsertProduct, err)
		return output, err
	}

	output.Id = input.Id

	return output, nil
}

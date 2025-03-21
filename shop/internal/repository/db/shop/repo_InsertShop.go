package shop

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
)

func (r *RepoDbShop) InsertShop(ctx context.Context, input model.InsertShopInput) (output model.InsertShopOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableShop).
		Set("id", input.Id).
		Set("owner_id", input.OwnerId).
		Set("name", input.Name).
		Set("status", input.Status).
		Set("created_at", input.CreatedAt).
		Set("created_by", input.CreatedBy)

	_, err = stmt.ExecAndClose(ctx, r.db)
	if err != nil {
		err = errors.Join(in_err.ErrInsertShop, err)
		return output, err
	}

	output.Id = input.Id

	return output, nil
}

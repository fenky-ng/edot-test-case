package shop

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/shop/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/shop/internal/error"
	"github.com/fenky-ng/edot-test-case/shop/internal/model"
	"github.com/lib/pq"
)

func (r *RepoDbShop) GetShops(ctx context.Context, input model.GetShopsInput) (output model.GetShopsOutput, err error) {
	var item model.Shop

	stmt := r.sql.From(constant.TableShop).
		Select("id").To(&item.Id).
		Select("owner_id").To(&item.OwnerId).
		Select("name").To(&item.Name).
		Select("status").To(&item.Status).
		Where("deleted_at IS NULL").
		OrderBy("name ASC")

	if len(input.Ids) != 0 {
		stmt.Where("id = ANY(?)", pq.Array(input.Ids))
	}

	err = stmt.QueryAndClose(ctx, r.db, func(rows *sql.Rows) {
		output.Shops = append(output.Shops, item)
		item = model.Shop{} // reset
	})
	if err != nil {
		err = errors.Join(in_err.ErrGetShops, err)
		return output, err
	}

	return output, nil
}

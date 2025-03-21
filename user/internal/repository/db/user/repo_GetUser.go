package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	"github.com/google/uuid"
)

func (r *RepoDbUser) GetUser(ctx context.Context, input model.GetUserInput) (output model.GetUserOutput, err error) {
	stmt := r.sql.From(constant.TableUser).
		Select("id").To(&output.Id).
		Select("name").To(&output.Name).
		Select("COALESCE(phone, '') AS phone").To(&output.Phone).
		Select("COALESCE(email, '') AS email").To(&output.Email).
		Select("hashed_password").To(&output.HashedPassword).
		Select("COALESCE(deleted_at, 0) AS deleted_at").To(&output.DeletedAt)

	if input.Id != uuid.Nil {
		stmt.Where("id = ?", input.Id)
	} else if input.Phone != "" {
		stmt.Where("phone = ?", input.Phone)
	} else if input.Email != "" {
		stmt.Where("email = ?", input.Email)
	}

	err = stmt.QueryRowAndClose(ctx, r.db)
	if err != nil {
		appErr := in_err.ErrGetUser
		if errors.Is(err, sql.ErrNoRows) {
			appErr = in_err.ErrUserNotFound
		}
		err = errors.Join(appErr, err)
		return output, err
	}

	return output, nil
}

package user

import (
	"context"
	"errors"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
)

func (r *RepoDbUser) InsertUser(ctx context.Context, input model.InsertUserInput) (output model.InsertUserOutput, err error) {
	stmt := r.sql.InsertInto(constant.TableUser).
		Set("id", input.Id).
		Set("name", input.Name).
		Set("phone", input.Phone).
		Set("email", input.Email).
		Set("hashed_password", input.HashedPassword).
		Set("created_at", input.CreatedAt).
		Set("created_by", input.CreatedBy)

	_, err = stmt.ExecAndClose(ctx, r.db)
	if err != nil {
		err = errors.Join(in_err.ErrInsertUser, err)
		return output, err
	}

	output.Id = input.Id

	return output, nil
}

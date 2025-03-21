package user

import (
	"context"

	"github.com/fenky-ng/edot-test-case/user/internal/model"
)

type RepoDbUserInterface interface {
	GetUser(ctx context.Context, input model.GetUserInput) (output model.GetUserOutput, err error)
	InsertUser(ctx context.Context, input model.InsertUserInput) (output model.InsertUserOutput, err error)
}

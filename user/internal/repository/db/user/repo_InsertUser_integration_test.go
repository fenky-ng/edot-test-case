package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIntegration_RepoDbUser_InsertUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	db := test_util.NewLocalDb()
	r := db_user.InitRepoDbUser(db_user.InitRepoDbUserOptions{
		DB: db,
	})

	t.Run("should insert user successfully", func(t *testing.T) {
		// GIVEN
		// request
		ctx := context.Background()
		input := model.InsertUserInput{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          faker.Phonenumber(),
			Email:          faker.Email(),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			CreatedAt:      time.Now().UnixMilli(),
			CreatedBy:      constant.TestDataCreatedBy,
		}

		// user data
		defer deleteUserData(ctx, db, t, []tableUserData{
			{
				Id: input.Id,
			},
		})

		// WHEN
		gotRes, gotErr := r.InsertUser(ctx, input)

		// THEN
		require.NoError(t, gotErr)
		require.Equal(t, model.InsertUserOutput{
			Id: input.Id,
		}, gotRes)
	})
}

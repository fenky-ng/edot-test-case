package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	pointer_util "github.com/fenky-ng/edot-test-case/user/internal/utility/pointer"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestIntegration_RepoDbUser_GetUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	db := test_util.NewLocalDb()
	r := db_user.InitRepoDbUser(db_user.InitRepoDbUserOptions{
		DB: db,
	})

	currentUnixMilli := time.Now().UnixMilli()

	t.Run("should get user by id successfully", func(t *testing.T) {
		// GIVEN
		// user data
		user1 := tableUserData{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          pointer_util.PointerOf(faker.Phonenumber()),
			Email:          pointer_util.PointerOf(faker.Email()),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			CreatedAt:      currentUnixMilli,
			CreatedBy:      constant.TestDataCreatedBy,
			DeletedAt:      pointer_util.PointerOf(currentUnixMilli),
		}
		userData := []tableUserData{
			user1,
		}
		userDataCleanup := insertUserData(context.Background(), db, t, userData)
		defer userDataCleanup()

		// request
		ctx := context.Background()
		input := model.GetUserInput{
			Id: user1.Id,
		}

		// WHEN
		gotRes, gotErr := r.GetUser(ctx, input)

		// THEN
		require.NoError(t, gotErr)
		require.Equal(t, model.GetUserOutput{
			Id:             user1.Id,
			Name:           user1.Name,
			Phone:          pointer_util.ValueOf(user1.Phone),
			Email:          pointer_util.ValueOf(user1.Email),
			HashedPassword: user1.HashedPassword,
			DeletedAt:      pointer_util.ValueOf(user1.DeletedAt),
		}, gotRes)
	})

	t.Run("should get user by phone successfully", func(t *testing.T) {
		// GIVEN
		// user data
		user1 := tableUserData{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          pointer_util.PointerOf(faker.Phonenumber()),
			Email:          pointer_util.PointerOf(faker.Email()),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			CreatedAt:      currentUnixMilli,
			CreatedBy:      constant.TestDataCreatedBy,
			DeletedAt:      pointer_util.PointerOf(currentUnixMilli),
		}
		userData := []tableUserData{
			user1,
		}
		userDataCleanup := insertUserData(context.Background(), db, t, userData)
		defer userDataCleanup()

		// request
		ctx := context.Background()
		input := model.GetUserInput{
			Phone: pointer_util.ValueOf(user1.Phone),
		}

		// WHEN
		gotRes, gotErr := r.GetUser(ctx, input)

		// THEN
		require.NoError(t, gotErr)
		require.Equal(t, model.GetUserOutput{
			Id:             user1.Id,
			Name:           user1.Name,
			Phone:          pointer_util.ValueOf(user1.Phone),
			Email:          pointer_util.ValueOf(user1.Email),
			HashedPassword: user1.HashedPassword,
			DeletedAt:      pointer_util.ValueOf(user1.DeletedAt),
		}, gotRes)
	})

	t.Run("should get user by email successfully", func(t *testing.T) {
		// GIVEN
		// user data
		user1 := tableUserData{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          pointer_util.PointerOf(faker.Phonenumber()),
			Email:          pointer_util.PointerOf(faker.Email()),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			CreatedAt:      currentUnixMilli,
			CreatedBy:      constant.TestDataCreatedBy,
			DeletedAt:      pointer_util.PointerOf(currentUnixMilli),
		}
		userData := []tableUserData{
			user1,
		}
		userDataCleanup := insertUserData(context.Background(), db, t, userData)
		defer userDataCleanup()

		// request
		ctx := context.Background()
		input := model.GetUserInput{
			Email: pointer_util.ValueOf(user1.Email),
		}

		// WHEN
		gotRes, gotErr := r.GetUser(ctx, input)

		// THEN
		require.NoError(t, gotErr)
		require.Equal(t, model.GetUserOutput{
			Id:             user1.Id,
			Name:           user1.Name,
			Phone:          pointer_util.ValueOf(user1.Phone),
			Email:          pointer_util.ValueOf(user1.Email),
			HashedPassword: user1.HashedPassword,
			DeletedAt:      pointer_util.ValueOf(user1.DeletedAt),
		}, gotRes)
	})
}

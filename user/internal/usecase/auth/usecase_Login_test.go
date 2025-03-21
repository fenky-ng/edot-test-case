package auth_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/auth"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func Test_AuthUsecase_Login(t *testing.T) {
	type fields struct {
		RepoDbUser db_user.RepoDbUserInterface
	}
	type args struct {
		ctx   context.Context
		input model.LoginInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedJwt bool
		assertErr   require.ErrorAssertionFunc
		mock        func(tt *test)
	}

	var (
		ctx        = context.Background()
		phoneInput = model.LoginInput{
			Phone:    faker.Phonenumber(),
			Password: faker.Password(),
		}
		emailInput = model.LoginInput{
			Email:    faker.Email(),
			Password: faker.Password(),
		}
	)
	tests := []test{
		{
			name: "should return error if error occurred in repoDbUser.GetUser",
			args: args{
				ctx:   ctx,
				input: phoneInput,
			},
			expectedJwt: false,
			assertErr:   test_util.RequireErrorIs(in_err.ErrGetUser),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{},
					errors.Join(in_err.ErrGetUser, errors.New("expected GetUser error")),
				).Times(1)
			},
		},
		{
			name: "should return error if user not found by phone",
			args: args{
				ctx:   ctx,
				input: phoneInput,
			},
			expectedJwt: false,
			assertErr:   test_util.RequireErrorIs(in_err.ErrInvalidPhoneLogin),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{},
					errors.Join(in_err.ErrUserNotFound, sql.ErrNoRows),
				).Times(1)
			},
		},
		{
			name: "should return error if user not found by email",
			args: args{
				ctx:   ctx,
				input: emailInput,
			},
			expectedJwt: false,
			assertErr:   test_util.RequireErrorIs(in_err.ErrInvalidEmailLogin),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{},
					errors.Join(in_err.ErrUserNotFound, sql.ErrNoRows),
				).Times(1)
			},
		},
		{
			name: "should return error if password is incorrect",
			args: args{
				ctx:   ctx,
				input: phoneInput,
			},
			expectedJwt: false,
			assertErr:   test_util.RequireErrorIs(in_err.ErrInvalidPhoneLogin),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{
						Id:             uuid.MustParse(faker.UUIDHyphenated()),
						HashedPassword: hash_util.HashPassword(faker.Password()),
						DeletedAt:      time.Now().UnixMilli(),
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if account already deactivated",
			args: args{
				ctx:   ctx,
				input: phoneInput,
			},
			expectedJwt: false,
			assertErr:   test_util.RequireErrorIs(in_err.ErrUserDeactivated),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{
						Id:             uuid.MustParse(faker.UUIDHyphenated()),
						HashedPassword: hash_util.HashPassword(tt.args.input.Password),
						DeletedAt:      time.Now().UnixMilli(),
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return jwt if login was successful",
			args: args{
				ctx:   ctx,
				input: phoneInput,
			},
			expectedJwt: true,
			assertErr:   require.NoError,
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Phone: tt.args.input.Phone,
						Email: tt.args.input.Email,
					},
				).Return(
					model.GetUserOutput{
						Id:             uuid.MustParse(faker.UUIDHyphenated()),
						HashedPassword: hash_util.HashPassword(tt.args.input.Password),
					},
					nil,
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock(&tt)
			}

			u := auth.InitAuthUsecase(auth.InitAuthUsecaseOptions{
				RepoDbUser: tt.fields.RepoDbUser,
			})
			gotRes, gotErr := u.Login(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			if tt.expectedJwt {
				require.NotEmpty(t, gotRes.JWT)
			}
		})
	}
}

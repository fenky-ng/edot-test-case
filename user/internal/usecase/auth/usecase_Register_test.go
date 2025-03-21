package auth_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/auth"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_AuthUsecase_Register(t *testing.T) {
	type fields struct {
		RepoDbUser db_user.RepoDbUserInterface
	}
	type args struct {
		ctx   context.Context
		input model.RegisterInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.RegisterOutput
		assertErr   require.ErrorAssertionFunc
		mock        func(tt *test)
	}

	var (
		ctx   = context.Background()
		input = model.RegisterInput{
			Name:     faker.Name(),
			Phone:    faker.Phonenumber(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		expectedValidOutput = model.RegisterOutput{
			Id: uuid.MustParse(faker.UUIDHyphenated()),
		}
	)
	tests := []test{
		{
			name: "should return error if error occurred in repoDbUser.GetUser",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.RegisterOutput{},
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
			name: "should return error if phone already registered",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.RegisterOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrPhoneRegistered),
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
						Id:    uuid.MustParse(faker.UUIDHyphenated()),
						Phone: tt.args.input.Phone,
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if email already registered",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.RegisterOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrEmailRegistered),
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
						Id:    uuid.MustParse(faker.UUIDHyphenated()),
						Email: tt.args.input.Email,
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if account already deactivated",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.RegisterOutput{},
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
						Id:        uuid.MustParse(faker.UUIDHyphenated()),
						DeletedAt: time.Now().UnixMilli(),
					},
					nil,
				).Times(1)
			},
		},
		{
			name: "should return error if error occurred in repoDbUser.InsertUser",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.RegisterOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrInsertUser),
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

				repoDbUserMock.EXPECT().InsertUser(
					tt.args.ctx,
					model.InsertUserInput{
						Name:           tt.args.input.Name,
						Phone:          tt.args.input.Phone,
						Email:          tt.args.input.Email,
						HashedPassword: hash_util.HashPassword(tt.args.input.Password),
						CreatedBy:      constant.ServiceName,
					}.Matcher(),
				).Return(
					model.InsertUserOutput{},
					errors.Join(in_err.ErrInsertUser, errors.New("expected InsertUser error")),
				).Times(1)
			},
		},
		{
			name: "should return id of the user if registration was successful",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: expectedValidOutput,
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
					model.GetUserOutput{},
					errors.Join(in_err.ErrUserNotFound, sql.ErrNoRows),
				).Times(1)

				repoDbUserMock.EXPECT().InsertUser(
					tt.args.ctx,
					model.InsertUserInput{
						Name:           tt.args.input.Name,
						Phone:          tt.args.input.Phone,
						Email:          tt.args.input.Email,
						HashedPassword: hash_util.HashPassword(tt.args.input.Password),
						CreatedBy:      constant.ServiceName,
					}.Matcher(),
				).Return(
					model.InsertUserOutput{
						Id: tt.expectedRes.Id,
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
			gotRes, gotErr := u.Register(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

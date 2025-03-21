package profile_test

import (
	"context"
	"errors"
	"testing"

	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	"github.com/fenky-ng/edot-test-case/user/internal/usecase/profile"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_ProfileUsecase_GetProfile(t *testing.T) {
	type fields struct {
		RepoDbUser db_user.RepoDbUserInterface
	}
	type args struct {
		ctx   context.Context
		input model.GetProfileInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.GetProfileOutput
		assertErr   require.ErrorAssertionFunc
		mock        func(tt *test)
	}

	var (
		ctx   = context.Background()
		input = model.GetProfileInput{
			Id: uuid.New(),
		}

		expectedValidOutput = model.GetProfileOutput{
			Id:    uuid.MustParse(faker.UUIDHyphenated()),
			Name:  faker.Name(),
			Phone: faker.Phonenumber(),
			Email: faker.Email(),
		}
	)
	tests := []test{
		{
			name: "should return error if error occurred in repoDbUser.GetUser",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.GetProfileOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrGetUser),
			mock: func(tt *test) {
				mockCtrl := gomock.NewController(t)

				repoDbUserMock := db_user.NewMockRepoDbUserInterface(mockCtrl)

				tt.fields.RepoDbUser = repoDbUserMock

				repoDbUserMock.EXPECT().GetUser(
					tt.args.ctx,
					model.GetUserInput{
						Id: tt.args.input.Id,
					},
				).Return(
					model.GetUserOutput{},
					errors.Join(in_err.ErrGetUser, errors.New("expected GetUser error")),
				).Times(1)
			},
		},
		{
			name: "should return data if successfully get user",
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
						Id: tt.args.input.Id,
					},
				).Return(
					model.GetUserOutput{
						Id:    tt.expectedRes.Id,
						Name:  tt.expectedRes.Name,
						Phone: tt.expectedRes.Phone,
						Email: tt.expectedRes.Email,
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

			u := profile.InitProfileUsecase(profile.InitProfileUsecaseOptions{
				RepoDbUser: tt.fields.RepoDbUser,
			})
			gotRes, gotErr := u.GetProfile(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

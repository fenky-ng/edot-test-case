package user_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_RepoDbUser_GetUser(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx   context.Context
		input model.GetUserInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.GetUserOutput
		assertErr   require.ErrorAssertionFunc
		sqlMock     sqlmock.Sqlmock
		mock        func(tt *test)
	}

	var (
		expectedValidOutput = model.GetUserOutput{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          faker.Phonenumber(),
			Email:          faker.Email(),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			DeletedAt:      time.Now().UnixMilli(),
		}

		ctx           = context.Background()
		filterIdInput = model.GetUserInput{
			Id: expectedValidOutput.Id,
		}
		filterPhoneInput = model.GetUserInput{
			Phone: expectedValidOutput.Phone,
		}
		filterEmailInput = model.GetUserInput{
			Email: expectedValidOutput.Email,
		}
	)

	tests := []test{
		{
			name: "should return error get user if error occurred in query",
			args: args{
				ctx:   ctx,
				input: filterIdInput,
			},
			expectedRes: model.GetUserOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrGetUser),
			mock: func(tt *test) {
				tt.sqlMock.ExpectQuery(``).WithArgs(
					tt.args.input.Id,
				).WillReturnError(
					errors.New("expected error"),
				)
			},
		},
		{
			name: "should return error user not found if got error sql no rows",
			args: args{
				ctx:   ctx,
				input: filterIdInput,
			},
			expectedRes: model.GetUserOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrUserNotFound),
			mock: func(tt *test) {
				tt.sqlMock.ExpectQuery(``).WithArgs(
					tt.args.input.Id,
				).WillReturnError(
					sql.ErrNoRows,
				)
			},
		},
		{
			name: "should return data if user found by id",
			args: args{
				ctx:   ctx,
				input: filterIdInput,
			},
			expectedRes: expectedValidOutput,
			assertErr:   require.NoError,
			mock: func(tt *test) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "phone", "email", "hashed_password", "deleted_at",
				}).AddRow(
					tt.expectedRes.Id, tt.expectedRes.Name, tt.expectedRes.Phone,
					tt.expectedRes.Email, tt.expectedRes.HashedPassword, tt.expectedRes.DeletedAt,
				)
				tt.sqlMock.ExpectQuery(``).WithArgs(
					tt.args.input.Id,
				).WillReturnRows(
					rows,
				)
			},
		},
		{
			name: "should return data if user found by phone",
			args: args{
				ctx:   ctx,
				input: filterPhoneInput,
			},
			expectedRes: expectedValidOutput,
			assertErr:   require.NoError,
			mock: func(tt *test) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "phone", "email", "hashed_password", "deleted_at",
				}).AddRow(
					tt.expectedRes.Id, tt.expectedRes.Name, tt.expectedRes.Phone,
					tt.expectedRes.Email, tt.expectedRes.HashedPassword, tt.expectedRes.DeletedAt,
				)
				tt.sqlMock.ExpectQuery(``).WithArgs(
					tt.args.input.Phone,
				).WillReturnRows(
					rows,
				)
			},
		},
		{
			name: "should return data if user found by email",
			args: args{
				ctx:   ctx,
				input: filterEmailInput,
			},
			expectedRes: expectedValidOutput,
			assertErr:   require.NoError,
			mock: func(tt *test) {
				rows := sqlmock.NewRows([]string{
					"id", "name", "phone", "email", "hashed_password", "deleted_at",
				}).AddRow(
					tt.expectedRes.Id, tt.expectedRes.Name, tt.expectedRes.Phone,
					tt.expectedRes.Email, tt.expectedRes.HashedPassword, tt.expectedRes.DeletedAt,
				)
				tt.sqlMock.ExpectQuery(``).WithArgs(
					tt.args.input.Email,
				).WillReturnRows(
					rows,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock, sqlMock, err := sqlmock.New()
			require.NoError(t, err)
			defer dbMock.Close()

			tt.fields.DB = dbMock
			tt.sqlMock = sqlMock

			if tt.mock != nil {
				tt.mock(&tt)
			}

			r := db_user.InitRepoDbUser(db_user.InitRepoDbUserOptions{
				DB: tt.fields.DB,
			})
			gotRes, gotErr := r.GetUser(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

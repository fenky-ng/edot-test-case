package user_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	in_err "github.com/fenky-ng/edot-test-case/user/internal/error"
	"github.com/fenky-ng/edot-test-case/user/internal/model"
	db_user "github.com/fenky-ng/edot-test-case/user/internal/repository/db/user"
	hash_util "github.com/fenky-ng/edot-test-case/user/internal/utility/hash"
	test_util "github.com/fenky-ng/edot-test-case/user/internal/utility/test"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_RepoDbUser_InsertUser(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx   context.Context
		input model.InsertUserInput
	}
	type test struct {
		name        string
		fields      fields
		args        args
		expectedRes model.InsertUserOutput
		assertErr   require.ErrorAssertionFunc
		sqlMock     sqlmock.Sqlmock
		mock        func(tt *test)
	}

	var (
		ctx   = context.Background()
		input = model.InsertUserInput{
			Id:             uuid.New(),
			Name:           faker.Name(),
			Phone:          faker.Phonenumber(),
			Email:          faker.Email(),
			HashedPassword: hash_util.HashPassword(faker.Password()),
			CreatedAt:      time.Now().UnixMilli(),
			CreatedBy:      constant.ServiceName,
		}

		expectedValidOutput = model.InsertUserOutput{
			Id: input.Id,
		}
	)

	tests := []test{
		{
			name: "should return error insert user if error occurred in query",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: model.InsertUserOutput{},
			assertErr:   test_util.RequireErrorIs(in_err.ErrInsertUser),
			mock: func(tt *test) {
				tt.sqlMock.ExpectExec(``).WithArgs(
					tt.args.input.Id,
					tt.args.input.Name,
					tt.args.input.Phone,
					tt.args.input.Email,
					tt.args.input.HashedPassword,
					tt.args.input.CreatedAt,
					tt.args.input.CreatedBy,
				).WillReturnError(
					errors.New("expected error"),
				)
			},
		},
		{
			name: "should return id if successfully insert user",
			args: args{
				ctx:   ctx,
				input: input,
			},
			expectedRes: expectedValidOutput,
			assertErr:   require.NoError,
			mock: func(tt *test) {
				result := sqlmock.NewResult(0, 1)
				tt.sqlMock.ExpectExec(``).WithArgs(
					tt.args.input.Id,
					tt.args.input.Name,
					tt.args.input.Phone,
					tt.args.input.Email,
					tt.args.input.HashedPassword,
					tt.args.input.CreatedAt,
					tt.args.input.CreatedBy,
				).WillReturnResult(
					result,
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
			gotRes, gotErr := r.InsertUser(tt.args.ctx, tt.args.input)
			tt.assertErr(t, gotErr)
			require.Equal(t, tt.expectedRes, gotRes)
		})
	}
}

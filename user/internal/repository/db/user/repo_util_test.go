package user_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/fenky-ng/edot-test-case/user/internal/constant"
	"github.com/google/uuid"
	"github.com/leporo/sqlf"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type tableUserData struct {
	Id             uuid.UUID
	Name           string
	Phone          *string
	Email          *string
	HashedPassword string
	CreatedAt      int64
	CreatedBy      string
	UpdatedAt      *int64
	UpdatedBy      *string
	DeletedAt      *int64
}

func insertUserData(ctx context.Context, db *sql.DB, t *testing.T, data []tableUserData) (cleanup func()) {
	stmt := sqlf.PostgreSQL.InsertInto(constant.TableUser)
	for _, d := range data {
		stmt.NewRow().
			Set("id", d.Id).
			Set("name", d.Name).
			Set("phone", d.Phone).
			Set("email", d.Email).
			Set("hashed_password", d.HashedPassword).
			Set("created_at", d.CreatedAt).
			Set("created_by", d.CreatedBy).
			Set("updated_at", d.UpdatedAt).
			Set("updated_by", d.UpdatedBy).
			Set("deleted_at", d.DeletedAt)
	}
	if len(data) > 0 {
		_, err := stmt.ExecAndClose(ctx, db)
		require.NoError(t, err)
	}

	cleanup = func() {
		deleteUserData(ctx, db, t, data)
	}

	return
}

func deleteUserData(ctx context.Context, db *sql.DB, t *testing.T, data []tableUserData) {
	if len(data) > 0 {
		ids := make([]any, 0)
		for _, d := range data {
			ids = append(ids, d.Id)
		}

		_, err := sqlf.PostgreSQL.DeleteFrom(constant.TableUser).
			Where("id = ANY(?)", pq.Array(ids)).
			ExecAndClose(ctx, db)
		require.NoError(t, err)
	}
}

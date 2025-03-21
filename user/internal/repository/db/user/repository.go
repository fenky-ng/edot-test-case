package user

import (
	"database/sql"

	"github.com/leporo/sqlf"
)

type RepoDbUser struct {
	db  *sql.DB
	sql *sqlf.Dialect
}

type InitRepoDbUserOptions struct {
	DB *sql.DB
}

func InitRepoDbUser(opts InitRepoDbUserOptions) *RepoDbUser {
	return &RepoDbUser{
		db:  opts.DB,
		sql: sqlf.PostgreSQL,
	}
}

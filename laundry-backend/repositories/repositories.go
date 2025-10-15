package repositories

import (
	"context"
	"database/sql"
)

type Repositories struct {
	Db  *sql.DB
	Ctx context.Context
}

func NewRepositories(Db *sql.DB, Ctx context.Context) Repositories {
	return Repositories{
		Db:  Db,
		Ctx: Ctx,
	}
}
package postgres

import (
	"github.com/jackc/pgconn"
)

func UniqueKey(err error) bool {
	return err.(*pgconn.PgError).Code == "23505"
}

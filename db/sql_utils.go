package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/mattn/go-sqlite3"
)

func IsUniqueViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}

	var sqliteErr sqlite3.Error

	if errors.As(err, &sqliteErr) {
		return errors.Is(sqliteErr.Code, sqlite3.ErrConstraint)
	}

	return false
}

func IsNoRowsErr(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("Error is pg no rows")
		return err.Code == "02000"
	}

	return err == sql.ErrNoRows
}

package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(driver string, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package config

import (
	"database/sql"
)

// Env holds data about server environment
type Env struct {
	Connection *sql.DB
}

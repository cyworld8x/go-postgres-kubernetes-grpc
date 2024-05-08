package postgres

import (
	"database/sql"
	"time"
)

type Option func(*postgres)

type postgres struct {
	connAttempts int
	connTimeout  time.Duration

	db *sql.DB
}

type DBEngine interface {
	GetDB() *sql.DB
	Configure(...Option) DBEngine
	Close()
}

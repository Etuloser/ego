package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultConnMaxLifetime = time.Minute * 3
	_defaultMaxOpenConns    = 10
	_defaultMaxIdleConns    = 10
)

type Client struct {
	DB              *sql.DB
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

// New -.
func New(dsn string, opts ...Option) (*Client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	client := &Client{
		DB:              db,
		ConnMaxLifetime: _defaultConnMaxLifetime,
		MaxOpenConns:    _defaultMaxOpenConns,
		MaxIdleConns:    _defaultMaxIdleConns,
	}
	// Custom options
	for _, opt := range opts {
		opt(client)
	}
	return client, nil
}

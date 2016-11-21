package skt

import "database/sql"
import "time"

// DB sql.DB
type DB struct {
	*sql.DB
}

// CK sql.CoreKey
type CK struct {
	ID      int64     `db:"integer primary key"`
	Created time.Time `db:"datetime"`
	Updated time.Time `db:"datetime"`
}

// Open sql.Open
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

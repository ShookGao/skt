package skt

import (
	"database/sql"
	"time"
)

const (
	// DATE 日期格式
	DATE = "2006-01-02"
	// DATETIMEH 日期时间精确到小时
	DATETIMEH = "2006-01-02 15"
	// DATETIMEM 日期时间精确到分钟
	DATETIMEM = "2006-01-02 15:04"
	// DATETIME 日期时间精确到秒
	DATETIME = "2006-01-02 15:04:05"
	// LONGTIME 日期时间精确到毫秒
	LONGTIME = "2006-01-02 15:04:05.999"
)

// DB sql.DB
type DB struct {
	*sql.DB
}

// TX sql.Tx
type Tx struct {
	*sql.Tx
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

// Go skt.tx
func Go(tx *sql.Tx) *Tx {
	return &Tx{tx}
}

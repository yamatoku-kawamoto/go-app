package database

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type DB struct {
	*bun.DB
}

func Open(config Config) (*DB, error) {
	switch config := config.(type) {
	case SqliteConfig, *SqliteConfig:
		conn := config.ConnectionString()
		sqldb, err := sql.Open(sqliteshim.ShimName, conn)
		if err != nil {
			return nil, err
		}
		db := &DB{
			bun.NewDB(sqldb, sqlitedialect.New()),
		}
		return db, nil
	}
	return nil, fmt.Errorf("unknown database type: %T", config)
}

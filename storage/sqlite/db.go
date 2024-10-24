package sqlite

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewDB(DSN string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", DSN)
	if err != nil {
		return db, err
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}
	return db, nil
}

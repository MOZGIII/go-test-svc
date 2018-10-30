package mydb

import (
	"context"

	"github.com/jackc/pgx"
)

// DB implements database operations using pgx connection pool.
type DB struct {
	Conn *pgx.ConnPool
}

// SelectOne simulates a meaningful database operation.
func (d *DB) SelectOne(ctx context.Context) (int, error) {
	var one int
	err := d.Conn.QueryRowEx(ctx, `SELECT 1`, nil).Scan(&one)
	if err != nil {
		return 0, err
	}
	return one, nil
}

package postgre

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(dbURL string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), dbURL)
}

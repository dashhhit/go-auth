package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPool(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

package postgres

import (
	"auth-golang/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
)

const MAXCONNS = 5

func NewPoolConfig(cfg *config.Postgres) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout)
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = MAXCONNS
	return poolConfig, nil
}

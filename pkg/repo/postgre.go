package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/url"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New() *pgxpool.Pool {
	connStr :=
		fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
			"postgres",
			url.QueryEscape("db_user"),
			url.QueryEscape("pwd123"),
			"localhost",
			"5432",
			"db_test",
			5)
	ctx, _ := context.WithCancel(context.Background())

	poolConfig, _ := pgxpool.ParseConfig(connStr)
	poolConfig.MaxConns = 5

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	for i := 0; i < 5; i++ {
		go func(count int) {
			_, err = pool.Exec(ctx, ";")
			if err != nil {
				log.Fatalf("Ping failed: %v\n", err)
			}
			log.Println(count, "Query OK!")
		}(i)
	}

	log.Println("Connection OK!")
	return pool

}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

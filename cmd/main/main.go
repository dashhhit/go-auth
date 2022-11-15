package main

import (
	"auth-golang/config"
	"auth-golang/internal/app"
	_ "auth-golang/migrations"
)

func main() {
	cfg := &config.Config{}
	cfg.Host = "localhost"
	cfg.Username = "db_user"
	cfg.Password = "pwd123"
	cfg.Postgres.Port = "5432"
	cfg.DbName = "postgres"
	cfg.Timeout = 5

	app.RunApp(cfg)
}

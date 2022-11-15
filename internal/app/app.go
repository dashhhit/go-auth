package app

import (
	"auth-golang/config"
	router "auth-golang/internal/controller/http"
	"auth-golang/internal/migrations"
	"auth-golang/pkg/repo"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RunApp(cfg *config.Config) {

	app := fiber.New()
	router.SetupRoutes(app)

	poolConfig, _ := postgres.NewPoolConfig(&cfg.Postgres)

	_, err := postgres.NewPool(poolConfig)
	if err != nil {
		log.Fatalf("connect to database failed: %v\n", err)
	}

	migrations.SetupGoose(poolConfig.ConnString())

	app.Listen(":8080")
}

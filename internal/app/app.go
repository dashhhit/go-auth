package app

import (
	"auth-golang/config"
	router "auth-golang/internal/controller/http"
	"auth-golang/internal/migrations"
	"auth-golang/internal/usecase"
	"auth-golang/internal/usecase/repo"
	"auth-golang/pkg/repo"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RunApp(cfg *config.Config) {

	poolConfig, _ := postgres.NewPoolConfig(&cfg.Postgres)

	pg, err := postgres.NewPool(poolConfig)
	if err != nil {
		log.Fatalf("connect to database failed: %v\n", err)
	}

	userUseCase := usecase.New(repo.New(pg))

	app := fiber.New()

	router.SetupRoutes(app, userUseCase)

	migrations.SetupGoose(poolConfig.ConnString())

	app.Listen(":8080")
}

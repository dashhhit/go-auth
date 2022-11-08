package app

import (
	router "auth-golang/internal/controller/http"
	"auth-golang/pkg/repo"
	"github.com/gofiber/fiber/v2"
)

func RunApp() {

	app := fiber.New()
	router.SetupRoutes(app)
	repo.New()
	app.Listen(":8080")
}

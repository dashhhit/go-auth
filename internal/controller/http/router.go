package router

import (
	"auth-golang/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App, t usecase.UserHandlers) {
	app.Use(cors.New())

	user := app.Group("/user")
	newUsersRoute(user, t)
}

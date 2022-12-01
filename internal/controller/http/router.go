package router

import (
	"auth-golang/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App, user usecase.UserHandlers, login usecase.LoginHandlers) {
	app.Use(cors.New())

	userGroup := app.Group("/user")
	newUsersRoute(userGroup, user)

	auth := app.Group("/auth")
	newLoginRoute(auth, login)
}

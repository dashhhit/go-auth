package router

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	user := app.Group("/user")
	UsersRoute(user)
}

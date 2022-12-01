package router

import (
	"auth-golang/internal/entity"
	"auth-golang/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type loginRoutes struct {
	t usecase.LoginHandlers
}

func newLoginRoute(route fiber.Router, t usecase.LoginHandlers) {
	r := &loginRoutes{t}

	{
		route.Post("/login", r.login)
	}
}

func (r *loginRoutes) login(ctx *fiber.Ctx) error {
	input := &entity.LoginInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    400,
				"message": "invalid input",
			},
		})
	}

	jwt, err := r.t.GenerateJWT(input)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    500,
				"message": err.Error(),
			},
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": fiber.Map{
			"code": 200,
			"data": jwt,
		},
	})

}

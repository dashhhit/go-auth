package router

import (
	"auth-golang/internal/entity"
	"auth-golang/internal/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type userRoutes struct {
	t usecase.UserHandlers
}

func newUsersRoute(router fiber.Router, t usecase.UserHandlers) {
	r := &userRoutes{t: t}

	{
		router.Post("/", r.createUser)
		router.Get("/:id<int>", r.getUser)
	}
}

func (r *userRoutes) createUser(ctx *fiber.Ctx) error {
	user := entity.User{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    400,
				"message": "invalid input",
			},
		})
	}
	err := r.t.Create(context.Background(), user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    500,
				"message": err.Error(),
			},
		})
	}

	return ctx.SendStatus(200)
}

func (r *userRoutes) getUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := r.t.Get(context.Background(), id)
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
			"data": fiber.Map{
				"user": user,
			},
			"code": 200,
		},
	})
}

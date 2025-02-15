package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_BookHistory/shared/utils"
)

func (c *UserController) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.GetByID(ctx.Context(), id)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User found successfully",
		"data":    user,
	})
}

func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	users, err := c.userService.GetAll(ctx.Context())
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Users found successfully",
		"data":    users,
	})
}

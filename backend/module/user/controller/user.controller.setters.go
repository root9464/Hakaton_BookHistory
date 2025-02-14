package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Hakaton_Zalupa/module/user/dto"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (c *UserController) Create(ctx *fiber.Ctx) error {
	dto := new(user_dto.CreateDto)
	if err := ctx.BodyParser(dto); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.Create(ctx.Context(), dto); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
	})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.userService.Delete(ctx.Context(), id); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

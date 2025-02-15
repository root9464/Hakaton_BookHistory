package application_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (c *ApplicationController) GetAll(ctx *fiber.Ctx) error {
	applications, err := c.applicationServ.GetAll(ctx.Context())
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Applications found successfully",
		"data":    applications,
	})
}

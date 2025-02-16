package application_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_BookHistory/shared/utils"
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

func (c *ApplicationController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	application, err := c.applicationServ.GetByID(ctx.Context(), id)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Application found successfully",
		"data":    application,
	})
}

func (c *ApplicationController) GetByStatus(ctx *fiber.Ctx) error {
	status := ctx.Params("status")
	applications, err := c.applicationServ.GetByStatus(ctx.Context(), status)
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

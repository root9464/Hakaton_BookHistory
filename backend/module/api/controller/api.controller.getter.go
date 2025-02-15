package api_controller

import "github.com/gofiber/fiber/v2"

func (c *ApiController) GetRecord(ctx *fiber.Ctx) error {
	resp, err:= c.apiService.GetRecord(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": fiber.Map{
			"result": resp,
		},
	})
}

func (c *ApiController) GetRecordByID(ctx *fiber.Ctx) error {
	resp, err := c.apiService.GetRecordByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"result": resp,
		},
	})
}



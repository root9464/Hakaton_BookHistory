package api_controller

import "github.com/gofiber/fiber/v2"

func (c *ApiController) GetRecords(ctx *fiber.Ctx) error {
	resp, err:= c.apiService.GetRecords(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *ApiController) GetRecordByID(ctx *fiber.Ctx) error {
	resp, err := c.apiService.GetRecordByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (c *ApiController) GetFullRecordByID(ctx *fiber.Ctx) error {
	resp, err := c.apiService.GetFullRecordByID(ctx.Context(), ctx.Params("id"))
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


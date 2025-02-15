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

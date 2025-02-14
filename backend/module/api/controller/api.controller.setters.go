package api_controller

import (
	"github.com/gofiber/fiber/v2"
	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
)

func (c *ApiController) CreateRecord(ctx *fiber.Ctx) error {
	dto := new(api_dto.CreateFeatureRequest)

	if err := ctx.BodyParser(dto); err != nil {
		c.logger.Info("01")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	c.logger.Infof("CreateRecord: %+v", dto)

	id, err := c.apiService.CreateRecord(ctx.Context(), dto)
	if err != nil {
		c.logger.Info("02")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "record created successfully",
		"data": fiber.Map{
			"id": id,
		},
	})

	return nil
}

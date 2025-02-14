package api_controller

import (
	"github.com/gofiber/fiber/v2"
	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
	api_service "github.com/root9464/Hakaton_Zalupa/module/api/service"
)

func (c *ApiController) CreateRecord(ctx *fiber.Ctx) error {
	dto := new(api_dto.CreateFeatureRequest)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	c.logger.Infof("CreateRecord: %+v", dto)

	err := api_service.


	return nil
}
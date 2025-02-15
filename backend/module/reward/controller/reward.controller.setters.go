package reward_controller

import (
	"github.com/gofiber/fiber/v2"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	reward_dto "github.com/root9464/Hakaton_Zalupa/module/reward/dto"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (c *RewardController) Create(ctx *fiber.Ctx) error {
	reward := new(reward_dto.CreateRewardDto)
	if err := ctx.BodyParser(reward); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Failed to parse multipart form"})
	}

	files := file_dto.CreateManyFileDto{
		Files: form.File["files"],
	}

	if err := c.rewardServ.Create(ctx.Context(), reward, &files); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Reward created successfully",
	})
}

func (c *RewardController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.rewardServ.Delete(ctx.Context(), id); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Reward deleted successfully",
	})
}

package reward_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (c *RewardController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	reward, err := c.rewardServ.GetByID(ctx.Context(), id)
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Reward found successfully",
		"data":    reward,
	})
}

func (c *RewardController) GetAll(ctx *fiber.Ctx) error {
	rewards, err := c.rewardServ.GetAll(ctx.Context())
	if err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Rewards found successfully",
		"data":    rewards,
	})
}

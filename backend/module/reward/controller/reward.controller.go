package reward_controller

import (
	"github.com/gofiber/fiber/v2"
	reward_service "github.com/root9464/Hakaton_Zalupa/module/reward/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IRewardController = (*RewardController)(nil)

type IRewardController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
}

type RewardController struct {
	rewardServ reward_service.IRewardService
	logger     *logger.Logger
}

func NewRewardController(rewardServ reward_service.IRewardService, logger *logger.Logger) *RewardController {
	return &RewardController{
		rewardServ: rewardServ,
		logger:     logger,
	}
}

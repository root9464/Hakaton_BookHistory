package reward_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	file_service "github.com/root9464/Hakaton_BookHistory/module/file/service"
	reward_controller "github.com/root9464/Hakaton_BookHistory/module/reward/controller"
	reward_repository "github.com/root9464/Hakaton_BookHistory/module/reward/repository"
	reward_service "github.com/root9464/Hakaton_BookHistory/module/reward/service"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

type RewardModule struct {
	rewardRepo       reward_repository.IRewardRepository
	rewardServ       reward_service.IRewardService
	rewardController reward_controller.IRewardController

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	fileServ file_service.IFileService
}

func NewRewardModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	fileServ file_service.IFileService,
) *RewardModule {
	return &RewardModule{
		logger:    logger,
		validator: validator,
		db:        db,
		fileServ:  fileServ,
	}
}

func (m *RewardModule) RewardRepository() reward_repository.IRewardRepository {
	if m.rewardRepo == nil {
		m.rewardRepo = reward_repository.NewRewardRepository(m.logger, m.db)
	}
	return m.rewardRepo
}

func (m *RewardModule) RewardService() reward_service.IRewardService {
	if m.rewardServ == nil {
		m.rewardServ = reward_service.NewRewardService(m.RewardRepository(), m.logger, m.validator, m.fileServ)
	}
	return m.rewardServ
}

func (m *RewardModule) RewardController() reward_controller.IRewardController {
	if m.rewardController == nil {
		m.rewardController = reward_controller.NewRewardController(m.RewardService(), m.logger)
	}
	return m.rewardController
}

func (m *RewardModule) RewardRoutes(router fiber.Router) {
	reward := router.Group("/reward")
	reward.Post("/", m.RewardController().Create)
	reward.Get("/", m.RewardController().GetAll)
	reward.Get("/:id", m.RewardController().GetByID)
	reward.Delete("/:id", m.RewardController().Delete)
}

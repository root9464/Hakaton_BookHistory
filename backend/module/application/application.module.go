package application_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	application_controller "github.com/root9464/Hakaton_BookHistory/module/application/controller"
	application_repository "github.com/root9464/Hakaton_BookHistory/module/application/repository"
	application_service "github.com/root9464/Hakaton_BookHistory/module/application/service"
	file_service "github.com/root9464/Hakaton_BookHistory/module/file/service"
	reward_service "github.com/root9464/Hakaton_BookHistory/module/reward/service"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

type ApplicationModule struct {
	applicationRepo       application_repository.IApplicationRepository
	applicationServ       application_service.IApplicationService
	applicationController application_controller.IApplicationController

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	fileServ   file_service.IFileService
	rewardServ reward_service.IRewardService
}

func NewApplicationModule(logger *logger.Logger, db *gorm.DB, fileServ file_service.IFileService, rewardServ reward_service.IRewardService) *ApplicationModule {
	return &ApplicationModule{
		logger:    logger,
		validator: validator.New(),
		db:        db,

		fileServ:   fileServ,
		rewardServ: rewardServ,
	}
}

func (m *ApplicationModule) ApplicationRepository() application_repository.IApplicationRepository {
	if m.applicationRepo == nil {
		m.applicationRepo = application_repository.NewApplicationRepository(m.logger, m.db)
	}
	return m.applicationRepo
}

func (m *ApplicationModule) ApplicationService() application_service.IApplicationService {
	if m.applicationServ == nil {
		m.applicationServ = application_service.NewApplicationService(m.logger, m.ApplicationRepository(), m.validator, m.fileServ, m.rewardServ)
	}
	return m.applicationServ
}

func (m *ApplicationModule) ApplicationController() application_controller.IApplicationController {
	if m.applicationController == nil {
		m.applicationController = application_controller.NewApplicationController(m.ApplicationService(), m.logger)
	}
	return m.applicationController
}

func (m *ApplicationModule) ApplicationRoutes(router fiber.Router) {
	application := router.Group("/application")
	application.Post("/", m.ApplicationController().Create)
	application.Get("/", m.ApplicationController().GetAll)
	application.Put("/:id", m.ApplicationController().UpdateStatus)
	application.Post("/email", m.ApplicationController().SendEmail)
	application.Get("/:id", m.ApplicationController().GetByID)
	application.Put("/", m.ApplicationController().UpdateAll)
	application.Get("/status/:status", m.ApplicationController().GetByStatus)
}

package application_controller

import (
	"github.com/gofiber/fiber/v2"
	application_service "github.com/root9464/Hakaton_Zalupa/module/application/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IApplicationController = (*ApplicationController)(nil)

type IApplicationController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	UpdateStatus(ctx *fiber.Ctx) error
}

type ApplicationController struct {
	applicationServ application_service.IApplicationService
	logger          *logger.Logger
}

func NewApplicationController(applicationServ application_service.IApplicationService, logger *logger.Logger) *ApplicationController {
	return &ApplicationController{
		applicationServ: applicationServ,
		logger:          logger,
	}
}

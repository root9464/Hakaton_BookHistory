package user_controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	user_service "github.com/root9464/Hakaton_Zalupa/module/user/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IUserController = (*UserController)(nil)

type IUserController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type UserController struct {
	userService user_service.IUserService
	logger      *logger.Logger
	validator   *validator.Validate
}

func NewUserController(
	userService user_service.IUserService,
	logger *logger.Logger,
	validator *validator.Validate,
) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
		validator:   validator,
	}
}

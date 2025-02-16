package user_module

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwt_module "github.com/root9464/Hakaton_Zalupa/module/jwt"
	user_controller "github.com/root9464/Hakaton_Zalupa/module/user/controller"
	user_repository "github.com/root9464/Hakaton_Zalupa/module/user/repository"
	user_service "github.com/root9464/Hakaton_Zalupa/module/user/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
	"github.com/root9464/Hakaton_Zalupa/shared/middleware"
	"gorm.io/gorm"
)

type UserModule struct {
	userService    user_service.IUserService
	userController user_controller.IUserController
	userRepo       user_repository.IUserRepository

	logger    *logger.Logger
	validator *validator.Validate
	db        *gorm.DB

	jwtModule jwt_module.JwtModule
	publicKey string
}

func NewUserModule(
	logger *logger.Logger,
	validator *validator.Validate,
	db *gorm.DB,
	jwtModule jwt_module.JwtModule,
	publicKey string,
) *UserModule {
	return &UserModule{
		logger:    logger,
		validator: validator,
		db:        db,
		jwtModule: jwtModule,
		publicKey: publicKey,
	}
}

func (m *UserModule) UserRepository() user_repository.IUserRepository {
	if m.userRepo == nil {
		m.userRepo = user_repository.NewUserRepository(m.logger, m.db)
	}
	return m.userRepo
}

func (m *UserModule) UserService() user_service.IUserService {
	if m.userService == nil {
		m.userService = user_service.NewUserService(m.UserRepository(), m.logger, m.validator)
	}
	return m.userService
}

func (m *UserModule) UserController() user_controller.IUserController {
	if m.userController == nil {
		m.userController = user_controller.NewUserController(m.UserService(), m.logger, m.validator)
	}
	return m.userController
}

func (m *UserModule) UserRoutes(router fiber.Router) {
	middleware := middleware.NewMiddleware(m.logger, m.userRepo, m.jwtModule.JwtHelpers(), m.publicKey)

	// user := router.Group("/user", middleware.UserOnly())
	user := router.Group("/user")
	user.Post("/", m.UserController().Create)
	user.Get("/", m.UserController().GetAll)
	user.Get("/:id", m.UserController().GetById)
	user.Delete("/:id", middleware.CreatorOnly(), m.UserController().Delete)
}

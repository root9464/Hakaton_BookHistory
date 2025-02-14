package api_module

import (
	"github.com/gofiber/fiber"
	"github.com/root9464/Hakaton_Zalupa/config"
	api_controller "github.com/root9464/Hakaton_Zalupa/module/api/controller"
	api_service "github.com/root9464/Hakaton_Zalupa/module/api/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

type ApiModule struct {
	apiService    api_service.IApiService
	apiController api_controller.IApiController

	logger *logger.Logger
	config *config.Config
}

func NewApiModule(logger *logger.Logger, config *config.Config,
	apiService api_service.IApiService, apiController api_controller.ApiController) *ApiModule {
	return &ApiModule{
		logger:        logger,
		config:        config,
		apiService:    apiService,
		apiController: apiController,
	}
}

func (m *ApiModule) AuthService() api_service.IApiService {
	if m.apiService == nil {
		m.apiService = api_service.NewAuthService(m.logger, m.config, m.apiService)
	}
	return m.apiService
}

func (m *ApiModule) AuthController() api_controller.IAuthController {

	if m.apiController == nil {
		m.apiController = api_controller.NewAuthController(m.AuthService(), m.logger, m.config)
	}
	return m.apiController
}

func (m *ApiModule) AuthRoutes(router fiber.Router) {
	auth := router.Group("/api")
	auth.Post("/authorize", m.AuthController().Authorize)
	auth.Post("/refresh", m.AuthController().RefreshAccessToken)
	auth.Get("/jwt-ping", m.AuthController().JwtPing)
}

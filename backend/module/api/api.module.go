package api_module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_BookHistory/config"
	api_controller "github.com/root9464/Hakaton_BookHistory/module/api/controller"
	api_service "github.com/root9464/Hakaton_BookHistory/module/api/service"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
)

type ApiModule struct {
	apiService    api_service.IApiService
	apiController api_controller.IApiController

	logger *logger.Logger
	config *config.Config
}

func NewApiModule(logger *logger.Logger, config *config.Config,
) *ApiModule {
	return &ApiModule{
		logger: logger,
		config: config,
	}
}

func (m *ApiModule) ApiService() api_service.IApiService {
	if m.apiService == nil {
		m.apiService = api_service.NewApiService(m.logger, m.config)
	}
	return m.apiService
}

func (m *ApiModule) ApiController() api_controller.IApiController {
	if m.apiController == nil {
		m.apiController = api_controller.NewApiController(m.ApiService(), m.logger, m.config)
	}
	return m.apiController
}

func (m *ApiModule) AuthRoutes(router fiber.Router) {
	auth := router.Group("/exapi")

	auth.Get("/get", m.ApiController().GetRecords)
	auth.Get("/get/:id", m.ApiController().GetRecordByID)
	auth.Get("/get-full/:id", m.ApiController().GetFullRecordByID)

	auth.Post("/create", m.ApiController().CreateRecord)
	auth.Delete("/delete", m.ApiController().DeleteRecord)
	auth.Put("/update/:id", m.ApiController().UpadateRecord)

	auth.Post("/upload", m.ApiController().UploadAttachment)
	auth.Post("/attach/:id", m.ApiController().AttachingFile)

	auth.Delete("/delete/:record_id/:attachment_id", m.ApiController().DeleteAttachment)

	//художественный формат
	auth.Get("art", m.ApiController().GetArtInfo)

}

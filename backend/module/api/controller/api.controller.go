package api_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Hakaton_BookHistory/config"
	api_service "github.com/root9464/Hakaton_BookHistory/module/api/service"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
)

type IApiController interface {
	//Добавление записи в слой (ресурс)
	CreateRecord(ctx *fiber.Ctx) error

	// //Запрос информации по записи
	GetRecords(ctx *fiber.Ctx) error

	GetRecordByID(ctx *fiber.Ctx) error
	GetFullRecordByID(ctx *fiber.Ctx) error

	// //Удаление записи
	DeleteRecord(ctx *fiber.Ctx) error

	// //Изменение записи
	UpadateRecord(ctx *fiber.Ctx) error

	//Загрузка вложения для созданной записи
	UploadAttachment(ctx *fiber.Ctx) error

	// Прикрепление файла к записи из пункта 3
	AttachingFile(ctx *fiber.Ctx) error

	//Удаление вложения
	DeleteAttachment(ctx *fiber.Ctx) error

	//получение художественной информации
	GetArtInfo(ctx *fiber.Ctx) error
}

type ApiController struct {
	apiService api_service.IApiService

	logger *logger.Logger
	config *config.Config
}

func NewApiController(apiService api_service.IApiService, logger *logger.Logger, config *config.Config) *ApiController {
	return &ApiController{
		apiService: apiService,
		logger:     logger,
		config:     config,
	}
}

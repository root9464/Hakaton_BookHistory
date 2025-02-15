package application_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	application_dto "github.com/root9464/Hakaton_Zalupa/module/application/dto"
	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
	application_repository "github.com/root9464/Hakaton_Zalupa/module/application/repository"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	file_service "github.com/root9464/Hakaton_Zalupa/module/file/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IApplicationService = (*ApplicationService)(nil)

type IApplicationService interface {
	Create(ctx context.Context, application *application_dto.CreateApplicationDto, files *file_dto.CreateManyFileDto) error
	GetAll(ctx context.Context) ([]application_model.Application, error)
}

type ApplicationService struct {
	repo      application_repository.IApplicationRepository
	logger    *logger.Logger
	validator *validator.Validate

	fileServ file_service.IFileService
}

func NewApplicationService(logger *logger.Logger, repo application_repository.IApplicationRepository, vavalidator *validator.Validate, fileServ file_service.IFileService) *ApplicationService {
	return &ApplicationService{
		repo:      repo,
		logger:    logger,
		validator: vavalidator,
		fileServ:  fileServ,
	}
}

package application_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	application_dto "github.com/root9464/Hakaton_BookHistory/module/application/dto"
	application_model "github.com/root9464/Hakaton_BookHistory/module/application/model"
	application_repository "github.com/root9464/Hakaton_BookHistory/module/application/repository"
	file_dto "github.com/root9464/Hakaton_BookHistory/module/file/dto"
	file_service "github.com/root9464/Hakaton_BookHistory/module/file/service"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
)

var _ IApplicationService = (*ApplicationService)(nil)

type IApplicationService interface {
	Create(ctx context.Context, application *application_dto.CreateApplicationDto, files *file_dto.CreateManyFileDto) error
	GetAll(ctx context.Context) ([]application_model.Application, error)
	UpdateStatus(ctx context.Context, id string, dto *application_dto.UpdateStatus) error
	SendEmail(ctx context.Context, email *application_dto.Email) error
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

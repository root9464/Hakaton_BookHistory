package application_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	application_dto "github.com/root9464/Hakaton_Zalupa/module/application/dto"
	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (s *ApplicationService) Create(ctx context.Context, dto *application_dto.CreateApplicationDto, files *file_dto.CreateManyFileDto) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validation error: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	applicationModel, err := utils.ConvertDtoToEntity[application_model.Application](dto)
	if err != nil {
		s.logger.Errorf("error converting dto to entity: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if err := s.repo.Create(ctx, applicationModel); err != nil {
		s.logger.Errorf("error creating application: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if err := s.fileServ.CreateMany(ctx, files); err != nil {
		s.logger.Errorf("error creating files: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

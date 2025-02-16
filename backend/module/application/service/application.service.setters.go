package application_service

import (
	"context"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	application_dto "github.com/root9464/Hakaton_Zalupa/module/application/dto"
	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	file_model "github.com/root9464/Hakaton_Zalupa/module/file/model"
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

	names, err := s.fileServ.CreateMany(ctx, files)
	if err != nil {
		s.logger.Errorf("error creating files: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	for _, name := range names {
		applicationModel.Files = append(applicationModel.Files, file_model.File{
			Name: name,
		})
	}

	if err := s.repo.Create(ctx, applicationModel); err != nil {
		s.logger.Errorf("error creating application: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *ApplicationService) UpdateStatus(ctx context.Context, id string, dto *application_dto.UpdateStatus) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validation error: %v", err)
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := s.repo.Update(ctx, &application_model.Application{ID: id, Status: application_model.Status(dto.Status)}); err != nil {
		s.logger.Errorf("error updating application: %v", err)
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *ApplicationService) SendEmail(ctx context.Context, dto *application_dto.Email) error {
	// if err := s.validator.Struct(dto); err != nil {
	// 	s.logger.Warnf("validation error: %v", err)
	// 	return &fiber.Error{
	// 		Code:    400,
	// 		Message: err.Error(),
	// 	}
	// }

	host := "smtp.gmail.com"
	port := "587"
	to := []string{"ivanbatutin6002@mail.ru"}

	from := "fakeroot94@gmail.com"
	password := "memm cchg tscz gioi"

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, to, []byte(dto.Message))
	if err != nil {
		s.logger.Errorf("error sending email: %v", err)
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

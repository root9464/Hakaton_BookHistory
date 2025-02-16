package application_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	application_model "github.com/root9464/Hakaton_BookHistory/module/application/model"
)

func (s *ApplicationService) GetAll(ctx context.Context) ([]application_model.Application, error) {
	applications, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("error getting applications: %v", err)
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return applications, nil
}

func (s *ApplicationService) GetByID(ctx context.Context, id string) (*application_model.Application, error) {
	application, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Errorf("error getting application: %v", err)
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return application, nil
}

package application_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
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

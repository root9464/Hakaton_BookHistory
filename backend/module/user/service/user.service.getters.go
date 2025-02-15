package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
)

func (s *UserService) GetByID(ctx context.Context, id string) (*user_model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if user == nil {
		s.logger.Warn("user not found")
		return nil, &fiber.Error{
			Code:    404,
			Message: "User not found",
		}
	}

	return user, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]user_model.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if len(users) == 0 {
		s.logger.Warn("user not found")
		return nil, &fiber.Error{
			Code:    404,
			Message: "User not found",
		}
	}

	return users, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*user_model.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if user == nil {
		s.logger.Warn("user not found")
		return nil, &fiber.Error{
			Code:    404,
			Message: "User not found",
		}
	}

	return user, nil
}

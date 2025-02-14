package auth_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	auth_dto "github.com/root9464/Hakaton_Zalupa/module/auth/dto"
	user_dto "github.com/root9464/Hakaton_Zalupa/module/user/dto"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (s *authService) Authorize(ctx context.Context, dto *auth_dto.AutorizeDto) (*user_model.User, error) {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	existUser, err := s.userService.GetByEmail(ctx, dto.Email)
	if err != nil {
		s.logger.Warnf("get user error: %s", err.Error())
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if existUser == nil || !utils.ComparePassword(existUser.Password, dto.Password) {
		s.logger.Warn("password or email is incorrect")
		return nil, &fiber.Error{
			Code:    401,
			Message: "password or email is incorrect",
		}
	}

	return existUser, nil
}

func (s *authService) Register(ctx context.Context, dto *auth_dto.RegisterDto) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	userDto := user_dto.CreateDto{
		Email:      dto.Email,
		Password:   dto.Password,
		Name:       dto.Name,
		Surname:    dto.Surname,
		Patronymic: dto.Patronymic,
		Phone:      dto.Phone,
	}

	err := s.userService.Create(ctx, &userDto)
	if err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

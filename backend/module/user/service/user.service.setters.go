package user_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Hakaton_Zalupa/module/user/dto"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	"github.com/root9464/Hakaton_Zalupa/shared/utils"
)

func (s *UserService) Create(ctx context.Context, dto *user_dto.CreateType) error {
	if err := s.validator.Struct(dto); err != nil {
		s.logger.Warnf("validate error: %s", err.Error())
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	dto.Password = utils.GeneratePassword(dto.Password)
	userModel, err := utils.ConvertDtoToEntity[user_model.User](dto)
	if err != nil {
		s.logger.Warnf("convert dto to entity error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if err := s.repo.Create(ctx, userModel); err != nil {
		s.logger.Warnf("create user error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Warnf("delete user error: %s", err.Error())
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}

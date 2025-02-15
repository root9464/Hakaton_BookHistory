package reward_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	file_dto "github.com/root9464/Hakaton_BookHistory/module/file/dto"
	file_model "github.com/root9464/Hakaton_BookHistory/module/file/model"
	reward_dto "github.com/root9464/Hakaton_BookHistory/module/reward/dto"
	reward_model "github.com/root9464/Hakaton_BookHistory/module/reward/model"
	"github.com/root9464/Hakaton_BookHistory/shared/utils"
)

func (s *RewardService) Create(ctx context.Context, reward *reward_dto.CreateRewardDto, files *file_dto.CreateManyFileDto) error {
	if err := s.validator.Struct(reward); err != nil {
		s.logger.Warnf("validation error: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
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

	rewardModel, err := utils.ConvertDtoToEntity[reward_model.Reward](reward)
	if err != nil {
		s.logger.Errorf("error converting dto to entity: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if len(names) > 0 {
		rewardModel.Image = file_model.File{
			Name: names[0],
		}
	}

	if err := s.repo.Create(ctx, rewardModel); err != nil {
		s.logger.Errorf("error creating reward: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *RewardService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Errorf("error deleting reward: %v", err)
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return nil
}

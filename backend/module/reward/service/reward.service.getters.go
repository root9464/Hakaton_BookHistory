package reward_service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	reward_model "github.com/root9464/Hakaton_BookHistory/module/reward/model"
)

func (s *RewardService) GetAll(ctx context.Context) ([]reward_model.Reward, error) {
	rewards, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("error getting rewards: %v", err)
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if len(rewards) == 0 {
		return nil, &fiber.Error{
			Code:    400,
			Message: "rewards not found",
		}
	}
	return rewards, nil
}

func (s *RewardService) GetByID(ctx context.Context, id string) (*reward_model.Reward, error) {
	reward, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Errorf("error getting reward: %v", err)
		return nil, &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	if reward == nil {
		return nil, &fiber.Error{
			Code:    400,
			Message: "reward not found",
		}
	}

	return reward, nil
}

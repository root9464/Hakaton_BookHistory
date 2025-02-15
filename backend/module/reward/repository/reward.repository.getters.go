package reward_repository

import (
	"context"

	reward_model "github.com/root9464/Hakaton_BookHistory/module/reward/model"
)

func (r *RewardRepository) GetAll(ctx context.Context) ([]reward_model.Reward, error) {
	var rewards []reward_model.Reward
	if err := r.db.Preload("Image").Find(&rewards).Error; err != nil {
		r.logger.Errorf("Error getting rewards: %v", err)
		return nil, err
	}
	return rewards, nil
}

func (r *RewardRepository) GetByID(ctx context.Context, id string) (*reward_model.Reward, error) {
	var reward reward_model.Reward
	if err := r.db.Preload("Image").First(&reward, "id = ?", id).Error; err != nil {
		r.logger.Errorf("Error getting reward: %v", err)
		return nil, err
	}
	return &reward, nil
}

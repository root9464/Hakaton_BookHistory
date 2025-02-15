package reward_repository

import (
	"context"

	reward_model "github.com/root9464/Hakaton_Zalupa/module/reward/model"
)

func (r *RewardRepository) Create(ctx context.Context, reward *reward_model.Reward) error {
	r.logger.Info("Creating reward...")
	if err := r.db.Create(reward).Error; err != nil {
		r.logger.Errorf("Error creating reward: %v", err)
		return err
	}
	r.logger.Info("Reward created")
	return nil
}

func (r *RewardRepository) Delete(ctx context.Context, id string) error {
	r.logger.Info("Deleting reward...")
	if err := r.db.Delete(&reward_model.Reward{}, "id = ?", id).Error; err != nil {
		r.logger.Errorf("Error deleting reward: %v", err)
		return err
	}
	r.logger.Info("Reward deleted")
	return nil
}

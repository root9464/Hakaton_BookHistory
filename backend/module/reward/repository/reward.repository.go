package reward_repository

import (
	"context"

	reward_model "github.com/root9464/Hakaton_BookHistory/module/reward/model"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

var _ IRewardRepository = (*RewardRepository)(nil)

type IRewardRepository interface {
	Create(ctx context.Context, reward *reward_model.Reward) error
	GetAll(ctx context.Context) ([]reward_model.Reward, error)
	GetByID(ctx context.Context, id string) (*reward_model.Reward, error)
	Delete(ctx context.Context, id string) error
}

type RewardRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewRewardRepository(logger *logger.Logger, db *gorm.DB) *RewardRepository {
	return &RewardRepository{
		logger: logger,
		db:     db,
	}
}

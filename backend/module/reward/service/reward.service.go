package reward_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	file_dto "github.com/root9464/Hakaton_Zalupa/module/file/dto"
	file_service "github.com/root9464/Hakaton_Zalupa/module/file/service"
	reward_dto "github.com/root9464/Hakaton_Zalupa/module/reward/dto"
	reward_model "github.com/root9464/Hakaton_Zalupa/module/reward/model"
	reward_repository "github.com/root9464/Hakaton_Zalupa/module/reward/repository"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IRewardService = (*RewardService)(nil)

type IRewardService interface {
	Create(ctx context.Context, reward *reward_dto.CreateRewardDto, files *file_dto.CreateManyFileDto) error
	GetAll(ctx context.Context) ([]reward_model.Reward, error)
	GetByID(ctx context.Context, id string) (*reward_model.Reward, error)
	Delete(ctx context.Context, id string) error
}

type RewardService struct {
	repo      reward_repository.IRewardRepository
	logger    *logger.Logger
	validator *validator.Validate

	fileServ file_service.IFileService
}

func NewRewardService(
	repo reward_repository.IRewardRepository,
	logger *logger.Logger,
	validator *validator.Validate,
	fileServ file_service.IFileService,
) *RewardService {
	return &RewardService{
		repo:      repo,
		logger:    logger,
		validator: validator,
		fileServ:  fileServ,
	}
}

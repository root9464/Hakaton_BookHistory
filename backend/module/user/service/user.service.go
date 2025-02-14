package user_service

import (
	"context"

	"github.com/go-playground/validator/v10"
	user_dto "github.com/root9464/Hakaton_Zalupa/module/user/dto"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	user_repository "github.com/root9464/Hakaton_Zalupa/module/user/repository"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Create(ctx context.Context, dto *user_dto.CreateType) error
	Delete(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*user_model.User, error)
	GetAll(ctx context.Context) ([]user_model.User, error)
}
type UserService struct {
	repo      user_repository.IUserRepository
	logger    *logger.Logger
	validator *validator.Validate
}

func NewUserService(
	repo user_repository.IUserRepository,
	logger *logger.Logger,
	validator *validator.Validate,
) *UserService {
	return &UserService{
		repo:      repo,
		logger:    logger,
		validator: validator,
	}
}

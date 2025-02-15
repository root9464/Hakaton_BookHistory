package user_repository

import (
	"context"

	user_model "github.com/root9464/Hakaton_BookHistory/module/user/model"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

var _ IUserRepository = (*UserRepository)(nil)

type IUserRepository interface {
	Create(ctx context.Context, user *user_model.User) error
	GetAll(ctx context.Context) ([]user_model.User, error)
	GetByID(ctx context.Context, id string) (*user_model.User, error)
	Delete(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*user_model.User, error)
}

type UserRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewUserRepository(logger *logger.Logger, db *gorm.DB) *UserRepository {
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}

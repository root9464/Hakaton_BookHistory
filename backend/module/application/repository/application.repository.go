package application_repository

import (
	"context"

	application_model "github.com/root9464/Hakaton_BookHistory/module/application/model"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

var _ IApplicationRepository = (*ApplicationRepository)(nil)

type IApplicationRepository interface {
	Create(ctx context.Context, application *application_model.Application) error
	GetAll(ctx context.Context) ([]application_model.Application, error)
	Update(ctx context.Context, application *application_model.Application) error
}
type ApplicationRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewApplicationRepository(logger *logger.Logger, db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{
		logger: logger,
		db:     db,
	}
}

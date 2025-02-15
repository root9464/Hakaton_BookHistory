package file_repository

import (
	"context"

	file_model "github.com/root9464/Hakaton_BookHistory/module/file/model"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

var _ IFileRepository = (*FileRepository)(nil)

type IFileRepository interface {
	CreateMany(ctx context.Context, files []file_model.File) error
}

type FileRepository struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewFileRepository(logger *logger.Logger, db *gorm.DB) *FileRepository {
	return &FileRepository{
		logger: logger,
		db:     db,
	}
}

package file_repository

import (
	"context"

	file_model "github.com/root9464/Hakaton_BookHistory/module/file/model"
)

func (r *FileRepository) CreateMany(ctx context.Context, files []file_model.File) error {
	r.logger.Info("creating files...")
	if err := r.db.WithContext(ctx).Create(&files).Error; err != nil {
		r.logger.Errorf("error creating files: %v", err)
		return err
	}
	r.logger.Info("files created")
	return nil
}

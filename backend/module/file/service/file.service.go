package file_service

import (
	"context"

	file_dto "github.com/root9464/Hakaton_BookHistory/module/file/dto"
	file_repository "github.com/root9464/Hakaton_BookHistory/module/file/repository"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
)

var _ IFileService = (*FileService)(nil)

type IFileService interface {
	CreateMany(ctx context.Context, files *file_dto.CreateManyFileDto) ([]string, error)
}
type FileService struct {
	repo   file_repository.IFileRepository
	logger *logger.Logger
}

func NewFileService(logger *logger.Logger, repo file_repository.IFileRepository) *FileService {
	return &FileService{
		repo:   repo,
		logger: logger,
	}
}

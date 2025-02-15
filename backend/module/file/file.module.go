package file_module

import (
	file_repository "github.com/root9464/Hakaton_Zalupa/module/file/repository"
	file_service "github.com/root9464/Hakaton_Zalupa/module/file/service"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
	"gorm.io/gorm"
)

type FileModule struct {
	fileServ file_service.IFileService
	fileRepo file_repository.IFileRepository

	logger *logger.Logger
	db     *gorm.DB
}

func NewFileModule(logger *logger.Logger, db *gorm.DB) *FileModule {
	return &FileModule{
		logger: logger,
		db:     db,
	}
}

func (m *FileModule) FileService() file_service.IFileService {
	if m.fileServ == nil {
		m.fileServ = file_service.NewFileService(m.logger, m.FileRepository())
	}
	return m.fileServ
}

func (m *FileModule) FileRepository() file_repository.IFileRepository {
	if m.fileRepo == nil {
		m.fileRepo = file_repository.NewFileRepository(m.logger, m.db)
	}
	return m.fileRepo
}

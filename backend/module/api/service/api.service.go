package api_service

import (
	"context"
	"mime/multipart"

	"github.com/root9464/Hakaton_BookHistory/config"
	api_dto "github.com/root9464/Hakaton_BookHistory/module/api/dto"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
)

var _ IApiService = (*apiService)(nil)

type IApiService interface {
	CreateRecord(ctx context.Context, dto *api_dto.CreateFeatureRequest) (*api_dto.CreateFeatureResponse, error)
	GetRecord(ctx context.Context) ([]UserResponse, error)
	GetRecordByID(ctx context.Context, id string) (map[string]interface{}, error)
	DeleteRecord(ctx context.Context, ids []api_dto.CreateFeatureResponse) error
	UpdateRecord(ctx context.Context, id string, dto *api_dto.CreateFeatureRequest) error

	UploadAttachment(ctx context.Context, file *multipart.FileHeader, fileName string) ([]map[string]interface{}, error)
	AttachingFile(ctx context.Context, recordID string, attachmentRequest api_dto.AttachmentRequest) (int, error)
	DeleteAttachment(ctx context.Context, recordID string, attachmentID string) error
}

type apiService struct {
	logger *logger.Logger
	config *config.Config
}

func NewApiService(logger *logger.Logger, config *config.Config) *apiService {
	return &apiService{
		logger: logger,
		config: config,
	}
}

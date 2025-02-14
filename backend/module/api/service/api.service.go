package api_service

import (
	"context"

	"github.com/root9464/Hakaton_Zalupa/config"
	api_dto "github.com/root9464/Hakaton_Zalupa/module/api/dto"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
)

var _ IApiService = (*apiService)(nil)

type IApiService interface {
	CreateRecord(ctx context.Context, dto *api_dto.CreateFeatureRequest) (*api_dto.CreateFeatureResponse, error)
}

type apiService struct {
	logger *logger.Logger
	config *config.Config

	
}

func NewApiService(logger *logger.Logger, config *config.Config) *apiService {
	return &apiService{
		logger:     logger,
		config:     config,
	}
}


package application_repository

import (
	"context"

	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
)

func (r *ApplicationRepository) GetAll(ctx context.Context) ([]application_model.Application, error) {
	var applications []application_model.Application
	if err := r.db.Find(&applications).Error; err != nil {
		r.logger.Errorf("error getting applications: %v", err)
		return nil, err
	}
	return applications, nil
}

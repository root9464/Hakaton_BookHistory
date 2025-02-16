package application_repository

import (
	"context"

	application_model "github.com/root9464/Hakaton_BookHistory/module/application/model"
)

func (r *ApplicationRepository) GetAll(ctx context.Context) ([]application_model.Application, error) {
	var applications []application_model.Application
	if err := r.db.Preload("Files").Find(&applications).Error; err != nil {
		r.logger.Errorf("error getting applications: %v", err)
		return nil, err
	}
	return applications, nil
}

func (r *ApplicationRepository) GetByID(ctx context.Context, id string) (*application_model.Application, error) {
	var application application_model.Application
	if err := r.db.Preload("Files").First(&application, "id = ?", id).Error; err != nil {
		r.logger.Errorf("error getting application: %v", err)
		return nil, err
	}
	return &application, nil
}

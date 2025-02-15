package application_repository

import (
	"context"

	application_model "github.com/root9464/Hakaton_BookHistory/module/application/model"
)

func (r *ApplicationRepository) Create(ctx context.Context, applications *application_model.Application) error {
	r.logger.Info("creating application...")
	if err := r.db.Create(&applications).Error; err != nil {
		r.logger.Errorf("error creating application: %v", err)
		return err
	}
	r.logger.Info("application created")
	return nil
}

func (r *ApplicationRepository) Update(ctx context.Context, application *application_model.Application) error {
	r.logger.Info("updating application...")

	if err := r.db.Model(&application).Where("id = ?", application.ID).Updates(application).Error; err != nil {
		r.logger.Errorf("error updating application: %v", err)
		return err
	}

	r.logger.Info("application updated")
	return nil
}

func (r *ApplicationRepository) UpdateAll(ctx context.Context, application *application_model.Application) error {
	r.logger.Info("updating applications...")
	if err := r.db.Save(&application).Error; err != nil {
		r.logger.Errorf("error updating applications: %v", err)
		return err
	}
	r.logger.Info("applications updated")
	return nil
}

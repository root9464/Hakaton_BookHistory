package user_repository

import (
	"context"

	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
)

func (r *UserRepository) GetByID(ctx context.Context, id string) (*user_model.User, error) {
	r.logger.Info("Getting user by id...")
	var user user_model.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		r.logger.Errorf("Failed to get user by id: %v", err)
		return nil, err
	}
	r.logger.Info("User retrieved successfully")
	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]user_model.User, error) {
	r.logger.Info("Getting all users...")
	var users []user_model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		r.logger.Errorf("Failed to get all users: %v", err)
		return nil, err
	}
	r.logger.Info("Users retrieved successfully")
	return users, nil
}

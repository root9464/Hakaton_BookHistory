package user_repository

import (
	"context"

	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
)

func (r *UserRepository) Create(ctx context.Context, user *user_model.User) error {
	r.logger.Info("Creating user...")
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		r.logger.Errorf("Failed to create user: %v", err)
		return err
	}
	r.logger.Info("User created successfully")
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.logger.Info("Deleting user...")
	if err := r.db.WithContext(ctx).Delete(&user_model.User{}, id).Error; err != nil {
		r.logger.Errorf("Failed to delete user: %v", err)
		return err
	}
	r.logger.Info("User deleted successfully")
	return nil
}

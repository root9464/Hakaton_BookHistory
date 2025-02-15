package database

import (
	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
	file_model "github.com/root9464/Hakaton_Zalupa/module/file/model"
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
			&file_model.File{},
			&application_model.Application{},
		}

		log.Info("📦 Creating types...")

		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}

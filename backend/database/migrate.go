package database

import (
	file_model "github.com/root9464/Hakaton_BookHistory/module/file/model"
	user_model "github.com/root9464/Hakaton_BookHistory/module/user/model"
	"github.com/root9464/Hakaton_BookHistory/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
			&file_model.File{},
		}

		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user')")
		db.Exec("CREATE TYPE status AS ENUM('draft', 'published')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}

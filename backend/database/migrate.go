package database

import (
	user_model "github.com/root9464/Hakaton_Zalupa/module/user/model"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
		}

		log.Info("📦 Creating types...")

		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}

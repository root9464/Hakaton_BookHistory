package database

import (
	model "github.com/root9464/Hakaton_Zalupa/database/model"
	"github.com/root9464/Hakaton_Zalupa/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&model.User{},
			&model.Municipality{},
			&model.Soldier{},
			&model.Award{},
			&model.Document{},
			&model.Log{},
			&model.Template{},
			&model.SoldierTemplate{},
			//&model.Notification{},
		}

		log.Info("📦 Creating types...")

		db.Exec("CREATE TYPE selected_name AS ENUM('firstname', 'lastname', 'nickname', 'username')")
		db.Exec("CREATE TYPE role AS ENUM('administarator', 'user', 'creator', 'moderator')")
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		db.Exec("CREATE TYPE notification_type AS ENUM('info', 'event', 'invite', 'comment', 'message', 'like', 'dislike')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}

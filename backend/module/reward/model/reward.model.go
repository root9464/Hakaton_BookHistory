package reward_model

import file_model "github.com/root9464/Hakaton_Zalupa/module/file/model"

type Reward struct {
	ID    string          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name  string          `json:"name"`
	Image file_model.File `gorm:"polymorphic:Owner;" json:"image"`
}

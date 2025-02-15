package file_model

type File struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string `json:"name"`
	OwnerID   string `json:"owner_id"`   // ID владельца (Application или Reward)
	OwnerType string `json:"owner_type"` // Тип владельца (application или reward)
}

package file_model

type File struct {
	ID   string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string `json:"name"`
}

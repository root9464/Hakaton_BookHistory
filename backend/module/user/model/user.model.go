package user_model

import (
	"fmt"

	application_model "github.com/root9464/Hakaton_Zalupa/module/application/model"
)

type Role string

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type User struct {
	ID           string                          `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email        string                          `gorm:"unique;not null" json:"email"`
	Password     string                          `gorm:"not null" json:"password"`
	Name         string                          `gorm:"not null" json:"name"`
	Surname      string                          `gorm:"not null" json:"surname"`
	Patronymic   string                          `gorm:"not null" json:"patronymic"`
	Phone        string                          `gorm:"not null" json:"phone"`
	Role         Role                            `gorm:"column:role;type:role;not null;default:user" json:"role"`
	Applications []application_model.Application `gorm:"foreignKey:SenderID" json:"applications"`
}

func ParseRole(roleStr string) (*Role, error) {
	roles := map[string]Role{
		"user":           UserRole,
		"administarator": AdminRole,
	}

	if role, exists := roles[roleStr]; exists {
		return &role, nil
	}

	return nil, fmt.Errorf("invalid role: %s", roleStr)
}

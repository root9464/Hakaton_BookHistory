package application_model

import file_model "github.com/root9464/Hakaton_BookHistory/module/file/model"

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
)

type Application struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	SenderID string `json:"sender_id"`

	// Сведения о жителе Оренбургской области
	FIO                         string            `json:"fio"`
	Geom                        string            `json:"geom"`
	DateOfBirth                 string            `json:"date_of_birth"`
	PlaceOfBirth                string            `json:"place_of_birth"`
	NameOfMillitaryCommissariat string            `json:"name_of_millitary_commissariat"`
	MillitaryRank               string            `json:"millitary_rank"`
	DateOfDeath                 string            `json:"date_of_death"`
	BurialPlace                 string            `json:"burial_place"`
	BiographicalFacts           string            `json:"biographical_facts"`
	Status                      Status            `gorm:"column:status;type:status;not null;default:draft" json:"status"`
	Files                       []file_model.File `gorm:"polymorphic:Owner;" json:"files"`
}

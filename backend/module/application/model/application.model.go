package application_model

type Application struct {
	ID string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`

	SenderID string `json:"sender_id"`

	// Сведения о жителе Оренбургской области
	FIO                         string `json:"fio"`
	DateOfBirth                 string `json:"date_of_birth"`
	PlaceOfBirth                string `json:"place_of_birth"`
	NameOfMillitaryCommissariat string `json:"name_of_millitary_commissariat"`
	MillitaryRank               string `json:"millitary_rank"`
	DateOfDeath                 string `json:"date_of_death"`
	BurialPlace                 string `json:"burial_place"`
	BiographicalFacts           string `json:"biographical_facts"`
}

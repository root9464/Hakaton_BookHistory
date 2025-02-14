package database

import "time"

type User struct {
	ID           int       `json:"id"`
	PasswordHash string    `json:"-"`
	FullName     string    `json:"full_name"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Organization string    `json:"organization"`
	CreatedAt    time.Time `json:"created_at"`
}

type Municipality struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// type MilitaryConflict struct {
// 	ID          int       `json:"id"`
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

type Soldier struct {
	ID int `json:"id"`

	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	PatronicName string `json:"middle_name"`

	BirthDate  time.Time `json:"birth_date"`
	BirthPlace string    `json:"birth_place"`

	MilitaryCommissariat string `json:"military_commissariat"`
	Rank                 string `json:"rank"`

	//информация о вооруженных конфликтах, в которых принимал участие
	//житель Оренбургской области,
	Info string `json:"info"`

	DeathDate   time.Time `json:"death_date"`
	BurialPlace string    `json:"burial_place"`

	Biography string `json:"biography"`

	MunicipalityID int `json:"municipality_id"`
}

// type SoldierConflict struct {
// 	SoldierID  int `json:"soldier_id"`
// 	ConflictID int `json:"conflict_id"`
// }

type Award struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Keyname     string `json:"keyname"`
	Size        string `json:"size"`
	mime_type   string `json:"mime_type"`
	description string `json:"description"`
	is_image    bool   `json:"is_image"`
	file_meta   string `json:"file_meta"`
}

// type SoldierAward struct {
// 	SoldierID int `json:"soldier_id"`
// 	AwardID   int `json:"award_id"`
// }

type Document struct {
	ID          int       `json:"id"`
	SoldierID   int       `json:"soldier_id"`
	FileName    string    `json:"file_name"`
	FilePath    string    `json:"file_path"`
	FileType    string    `json:"file_type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Log struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	SoldierID int       `json:"soldier_id"`
	Action    string    `json:"action"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}

type Template struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type SoldierTemplate struct {
	SoldierID  int `json:"soldier_id"`
	TemplateID int `json:"template_id"`
}

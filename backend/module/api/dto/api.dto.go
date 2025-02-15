package api_dto

type Feature struct {
	ID         int    `json:"id,omitempty"` // Уникальный идентификатор записи
	Geom       string `json:"geom"`         // Геометрия точки в формате WKT (EPSG 3857)
	Fields     Fields `json:"fields"`       // Поля с данными
	Extensions struct {
		Description string       `json:"description,omitempty"`
		Attachment  []Attachment `json:"attachment,omitempty"`
	} `json:"extensions,omitempty"`
}

type Fields struct {
	Num      float64    `json:"num"`                // Номер записи
	Region   string `json:"n_raion"`            // Муниципальное образование
	FullName string `json:"fio"`                // ФИО
	Years    string `json:"years"`              // Годы жизни
	Info     string `json:"info,omitempty"`     // Биография
	Conflict string `json:"kontrakt,omitempty"` // Вооруженный конфликт
	Awards   string `json:"nagrads,omitempty"`  // Награды
}

type Attachment struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	MimeType string `json:"mime_type"`
}

// API запросы
type GetFeaturesResponse struct {
	Features []Feature `json:"features"`
}

type CreateFeatureRequest struct {
	Geom       string `json:"geom"`   // Координаты точки
	Fields     Fields `json:"fields"` // Данные о погибшем
	Extensions struct {
		Attachment  *Attachment `json:"attachment,omitempty"`
		Description string      `json:"description,omitempty"`
	} `json:"extensions,omitempty"`
}

type CreateFeatureResponse struct {
	ID int `json:"id"` // Уникальный ID новой записи
}

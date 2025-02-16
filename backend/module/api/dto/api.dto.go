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

type FeatureCord struct {
	ID     string `json:"id,omitempty"` // Уникальный идентификатор записи
	Coords struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"cords"`
}

type Fields struct {
	UUID     string `json:"uuid"`               // Номер записи
	Region   string `json:"n_raion"`            // Муниципальное образование
	FullName string `json:"fio"`                // ФИО
	Years    string `json:"years"`              // Годы жизни
	Info     string `json:"info,omitempty"`     // Биография
	Conflict string `json:"kontrakt,omitempty"` // Вооруженный конфликт
	Awards   string `json:"nagrads,omitempty"`  // Награды
}

type Attachment struct {
	ID       float64 `json:"id"`
	Name     string  `json:"name"`
	Size     int     `json:"size"`
	MimeType string  `json:"mime_type"`
}

// API запросы
type GetFeaturesResponse struct {
	Features []Feature `json:"features"`
}

type CreateFeatureRequest struct {
	Fields     Fields `json:"fields"` // Данные о погибшем
	Extensions struct {
		Attachment  *Attachment `json:"attachment,omitempty"`
		Description string      `json:"description,omitempty"`
	} `json:"extensions,omitempty"`
	Geom string `json:"geom"` // Координаты точки
}

type CreateFeatureResponse struct {
	ID int `json:"id"` // Уникальный ID новой записи
}

type AttachmentRequest struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	MimeType   string `json:"mime_type"`
	FileUpload struct {
		ID   string `json:"id"`
		Size int    `json:"size"`
	} `json:"file_upload"`
}

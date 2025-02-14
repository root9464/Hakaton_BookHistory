package auth_dto

type AutorizeDto struct {
	Email    string `json:"email" validate:"required,email"`    // Обязательное поле, должно быть валидным email
	Password string `json:"password" validate:"required,min=8"` // Обязательное поле, минимум 8 символов
}

type RegisterDto struct {
	Email      string `json:"email" validate:"required,email"`    // Обязательное поле, должно быть валидным email
	Password   string `json:"password" validate:"required,min=8"` // Обязательное поле, минимум 8 символов
	Name       string `json:"name" validate:"required"`           // Обязательное поле, только буквы
	Surname    string `json:"surname" validate:"required"`        // Обязательное поле, только буквы
	Patronymic string `json:"patronymic" validate:"omitempty"`    // Необязательное поле, только буквы (если указано)
	Phone      string `json:"phone" validate:"required,e164"`     // Обязательное поле, формат E.164 (например, +79123456789)
}

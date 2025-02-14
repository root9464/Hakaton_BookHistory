package auth_dto

type AutorizeDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
}

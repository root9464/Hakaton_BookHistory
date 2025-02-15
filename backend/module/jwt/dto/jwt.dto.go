package jwt_dto

import user_model "github.com/root9464/Hakaton_BookHistory/module/user/model"

type UserData struct {
	ID   string `json:"id" validate:"required"`
	Role string `json:"role" validate:"required"`
}

type UserJwtPayload struct {
	Iss  string          `json:"iss" validate:"required"`
	Sub  string          `json:"sub" validate:"required"`
	Iat  int64           `json:"iat" validate:"required"`
	Exp  int64           `json:"exp" validate:"required"`
	Role user_model.Role `json:"role" validate:"required"`
	Hash string          `json:"user_hash" validate:"required"`
}

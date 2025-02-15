package file_dto

import "mime/multipart"

type CreateManyFileDto struct {
	Files []*multipart.FileHeader `json:"file"`
}

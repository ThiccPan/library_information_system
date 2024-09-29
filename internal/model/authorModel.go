package model

import "github.com/thiccpan/library_information_system/internal/entity"

type AuthorResponse struct {
	Data *entity.Author
}

type AuthorsResponse struct {
	Data []entity.Author
}

type CreateAuthorRequest struct {
	Name string
	Bio  string
}

type UpdateAuthorRequest struct {
	Id   uint
	Name string
	Bio  string
}

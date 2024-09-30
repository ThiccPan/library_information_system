package model

import "mime/multipart"

type CreateBookRequest struct {
	Name      string `json:"name"`
	Stock     uint   `json:"stock"`
	Author_id uint   `json:"author_id"`
	Topic_id  uint   `json:"topic_id"`
}

type UpdateBookRequest struct {
	Id        uint
	Name      string `json:"name"`
	Stock     uint   `json:"stock"`
	Author_id uint   `json:"author_id"`
	Topic_id  uint   `json:"topic_id"`
	CoverImg  *multipart.FileHeader
}

type QueryBookRequest struct {
	Id             uint
	Query          string
	FilterTopicsId uint
	FilterUsersId  uint
}

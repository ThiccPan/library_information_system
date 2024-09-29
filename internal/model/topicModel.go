package model

import "github.com/thiccpan/library_information_system/internal/entity"

type TopicResponse struct {
	Data *entity.Topic
}

type TopicsResponse struct {
	Data []entity.Topic
}

type CreateTopicRequest struct {
	Name string
}

type UpdateTopicRequest struct {
	Id   uint
	Name string
}

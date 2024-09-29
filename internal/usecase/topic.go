package usecase

import (
	"context"
	"errors"
	"net/http"

	"github.com/thiccpan/library_information_system/internal/entity"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/repository"
	"gorm.io/gorm"
)

type TopicUsecase struct {
	Db         *gorm.DB
	Repository *repository.TopicRepository
}

func NewTopicUsecase(db *gorm.DB, repo *repository.TopicRepository) *TopicUsecase {
	return &TopicUsecase{
		Db:         db,
		Repository: repo,
	}
}

func (tu *TopicUsecase) Add(ctx context.Context, request *model.CreateTopicRequest) (*entity.Topic, *model.Response) {
	topic := &entity.Topic{
		Name: request.Name,
	}
	if err := tu.Repository.Add(tu.Db, topic); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to add topic"}
	}
	return topic, &model.Response{Code: http.StatusOK, Message: "successfully added new topic"}
}

func (au *TopicUsecase) Get(ctx context.Context) ([]entity.Topic, *model.Response) {
	tx := au.Db
	topics, err := au.Repository.GetAll(tx)
	if err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get topics", Err: err}
	}
	return topics, &model.Response{Code: http.StatusOK, Message: "successfully get topics"}
}

func (au *TopicUsecase) GetById(ctx context.Context, id uint) (*entity.Topic, *model.Response) {
	tx := au.Db
	topics := &entity.Topic{Id: id}
	if err := au.Repository.GetById(tx, topics); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to get topic", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get topic", Err: err}
	}
	return topics, &model.Response{Code: http.StatusOK, Message: "successfully get topic"}
}

func (au *TopicUsecase) Update(ctx context.Context, request *model.UpdateTopicRequest) (*entity.Topic, *model.Response) {
	tx := au.Db
	topic := &entity.Topic{
		Id:   request.Id,
		Name: request.Name,
	}
	if err := au.Repository.UpdateById(tx, topic); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to update topic", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to update topic", Err: err}
	}
	return topic, &model.Response{Code: http.StatusOK, Message: "successfully update new topic"}
}

func (au *TopicUsecase) Delete(ctx context.Context, id uint) (*entity.Topic, *model.Response) {
	tx := au.Db
	topic := &entity.Topic{Id: id}
	if err := au.Repository.Delete(tx, topic); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to delete topic", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to delete topic", Err: err}
	}
	return topic, &model.Response{Code: http.StatusOK, Message: "successfully delete topic"}
}

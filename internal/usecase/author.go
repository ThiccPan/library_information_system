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

type AuthorUsecase struct {
	Db         *gorm.DB
	Repository *repository.AuthorRepository
}

func NewAuthorUsecase(db *gorm.DB, repo *repository.AuthorRepository) *AuthorUsecase {
	return &AuthorUsecase{
		Db:         db,
		Repository: repo,
	}
}

func (au *AuthorUsecase) Add(ctx context.Context, request *model.CreateAuthorRequest) (*entity.Author, *model.Response) {
	author := &entity.Author{
		Name: request.Name,
		Bio:  request.Bio,
	}
	if err := au.Repository.Add(au.Db, author); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to add author"}
	}
	return author, &model.Response{Code: http.StatusOK, Message: "successfully added new author"}
}

func (au *AuthorUsecase) Get(ctx context.Context) ([]entity.Author, *model.Response) {
	tx := au.Db
	authors, err := au.Repository.GetAll(tx)
	if err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get authors", Err: err}
	}
	return authors, &model.Response{Code: http.StatusOK, Message: "successfully added new author"}
}

func (au *AuthorUsecase) GetById(ctx context.Context, id uint) (*entity.Author, *model.Response) {
	tx := au.Db
	authors := &entity.Author{Id: id}
	if err := au.Repository.GetById(tx, authors); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to get author", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get author", Err: err}
	}
	return authors, &model.Response{Code: http.StatusOK, Message: "successfully get author"}
}

func (au *AuthorUsecase) Update(ctx context.Context, request *model.UpdateAuthorRequest) (*entity.Author, *model.Response) {
	tx := au.Db
	authors := &entity.Author{
		Id:        request.Id,
		Name:      request.Name,
		Bio:       request.Bio,
	}
	if err := au.Repository.UpdateById(tx, authors); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to update author", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to update author", Err: err}
	}
	return authors, &model.Response{Code: http.StatusOK, Message: "successfully update new author"}
}

func (au *AuthorUsecase) Delete(ctx context.Context, id uint) (*entity.Author, *model.Response) {
	tx := au.Db
	authors := &entity.Author{Id: id}
	if err := au.Repository.Delete(tx, authors); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to delete author", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to delete author", Err: err}
	}
	return authors, &model.Response{Code: http.StatusOK, Message: "successfully delete author"}
}

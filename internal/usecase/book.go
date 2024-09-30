package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/thiccpan/library_information_system/internal/entity"
	"github.com/thiccpan/library_information_system/internal/model"
	"github.com/thiccpan/library_information_system/internal/repository"
	"gorm.io/gorm"
)

const bookCoverLocation = "resource/book_cover/"

type BookUsecase struct {
	Db         *gorm.DB
	Repository *repository.BookRepository
}

func NewBookUsecase(db *gorm.DB, repo *repository.BookRepository) *BookUsecase {
	return &BookUsecase{
		Db:         db,
		Repository: repo,
	}
}

func (tu *BookUsecase) Add(ctx context.Context, request *model.CreateBookRequest) (*entity.Book, *model.Response) {
	book := &entity.Book{
		Name:      request.Name,
		Stock:     request.Stock,
		Author_id: request.Author_id,
		Topic_id:  request.Topic_id,
	}

	if err := tu.Repository.Add(tu.Db, book); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to add book"}
	}
	return book, &model.Response{Code: http.StatusOK, Message: "successfully added new book"}
}

func (au *BookUsecase) Get(ctx context.Context, request *model.QueryBookRequest) ([]entity.Book, *model.Response) {
	tx := au.Db
	books, err := au.Repository.GetAll(tx)
	if err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get books", Err: err}
	}
	return books, &model.Response{Code: http.StatusOK, Message: "successfully get books"}
}

func (au *BookUsecase) GetById(ctx context.Context, request *model.QueryBookRequest) (*entity.Book, *model.Response) {
	tx := au.Db
	books := &entity.Book{Id: request.Id}
	if err := au.Repository.GetById(tx, books); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to get book", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to get book", Err: err}
	}
	return books, &model.Response{Code: http.StatusOK, Message: "successfully get book"}
}

func (au *BookUsecase) Update(ctx context.Context, request *model.UpdateBookRequest) (*entity.Book, *model.Response) {
	tx := au.Db.WithContext(ctx).Begin()

	book := &entity.Book{
		Id:        request.Id,
		Name:      request.Name,
		Stock:     request.Stock,
		Author_id: request.Author_id,
		Topic_id:  request.Topic_id,
	}
	if request.CoverImg != nil {
		book.Cover_img = fmt.Sprintf("%d-img.jpeg", request.Id)
	}

	// update book data
	if err := au.Repository.UpdateById(tx, book); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to update book", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to update book", Err: err}
	}

	// copy file to storage
	if request.CoverImg != nil {
		_, err := moveFile(fmt.Sprint(bookCoverLocation, book.Cover_img), request.CoverImg)
		if err != nil {
			tx.Rollback()
			return nil, &model.Response{Err: err, Code: http.StatusInternalServerError, Message: "failed to save cover image"}
		}
	}

	// commit the change
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &model.Response{Err: err, Code: http.StatusInternalServerError, Message: "failed to update book"}
	}

	return book, &model.Response{Code: http.StatusOK, Message: "successfully update book"}
}

func (au *BookUsecase) Delete(ctx context.Context, request *model.QueryBookRequest) (*entity.Book, *model.Response) {
	tx := au.Db
	book := &entity.Book{Id: request.Id}
	if err := au.Repository.Delete(tx, book); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to delete book", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to delete book", Err: err}
	}
	return book, &model.Response{Code: http.StatusOK, Message: "successfully delete book"}
}

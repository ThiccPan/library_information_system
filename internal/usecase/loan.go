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

type LoanUsecase struct {
	db         *gorm.DB
	repository *repository.LoanRepository
	bookRepo   *repository.BookRepository
}

func NewLoanUsecase(db *gorm.DB, repository *repository.LoanRepository, bookRepo *repository.BookRepository) *LoanUsecase {
	return &LoanUsecase{
		db:         db,
		repository: repository,
	}
}

func (lu *LoanUsecase) Add(c context.Context, req *model.CommandLoanRequest) (*entity.Loan, *model.Response) {
	tx := lu.db.WithContext(c).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	loan := &entity.Loan{
		User_id:       req.User_id,
		Book_id:       req.Book_id,
		LoanStatus_id: uint(entity.BORROWED_CODE),
		Deadline:      req.Deadline,
	}
	// fetch book
	book := &entity.Book{Id: loan.Book_id}
	if err := lu.bookRepo.GetById(tx, book); err != nil {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}
	// check if book is still available to loan
	if book.Stock < 1 {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusBadRequest, Err: errors.New("book is empty"), Message: "failed to add loan"}
	}
	// update book stock and update to data store
	book.Stock -= 1
	if err := lu.bookRepo.UpdateStock(tx, book); err != nil {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}

	// add loan
	if err := lu.repository.Add(tx, loan); err != nil {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}

	return loan, &model.Response{Code: http.StatusOK, Message: "adding new loan success"}
}

func (lu *LoanUsecase) Get(c context.Context, req *model.QueryLoanRequest) ([]entity.Loan, *model.Response) {
	tx := lu.db.WithContext(c)
	loans, err := lu.repository.GetAll(tx, req)
	if err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to get loans"}
	}

	return loans, &model.Response{Code: http.StatusOK, Message: "Successfully getting loans"}
}

func (lu *LoanUsecase) GetById(c context.Context, req *model.QueryLoanRequest) (*entity.Loan, *model.Response) {
	tx := lu.db.WithContext(c)
	loan := &entity.Loan{}
	if err := lu.repository.GetById(tx, loan); err != nil {
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to get loans"}
	}
	return loan, &model.Response{Code: http.StatusOK, Message: "Successfully getting loan"}
}

func (lu *LoanUsecase) Update(c context.Context, req *model.CommandLoanRequest) (*entity.Loan, *model.Response) {
	tx := lu.db.WithContext(c).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// get loan old data
	loan := &entity.Loan{Id: req.Id}
	if err := lu.repository.GetById(tx, loan); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to get loans",
		}
	}
	// check if book is returned already, if so then return error
	if loan.LoanStatus_id == uint(entity.RETURNED_CODE) {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusBadRequest, Message: "failed to update loan", Err: errors.New("loan is already returned")}
	}
	// if loan status is returned then update book stock
	if req.LoanStatus_id == uint(entity.RETURNED_CODE) {
		// get book from load id
		book := &entity.Book{Id: loan.Book_id}
		if err := lu.bookRepo.GetById(tx, book); err != nil {
			tx.Rollback()
			return nil, &model.Response{Code: http.StatusBadRequest, Message: "failed to update loan", Err: errors.New("failed to get updated book")}
		}
		// update book stock from fetched book
		book.Stock += 1
		if err := lu.bookRepo.UpdateById(tx, book); err != nil {
			tx.Rollback()
			return nil, &model.Response{Code: http.StatusBadRequest, Message: "failed to update loan", Err: errors.New("failed to update book stock")}
		}
	}

	// update loan
	loan.LoanStatus_id = req.LoanStatus_id
	loan.Book_id = req.Book_id
	loan.User_id = req.User_id
	loan.Deadline = req.Deadline
	if err := lu.repository.UpdateById(tx, loan); err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.Response{Code: http.StatusNotFound, Message: "failed to update book", Err: err}
		}
		return nil, &model.Response{Code: http.StatusInternalServerError, Message: "failed to update book", Err: err}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}

	return loan, &model.Response{Code: http.StatusOK, Message: "updating loan success"}
}

// delete returned loan
func (lu *LoanUsecase) Delete(c context.Context, req *model.CommandLoanRequest) (*entity.Loan, *model.Response) {
	tx := lu.db.WithContext(c).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	loan := &entity.Loan{Id: req.Id}
	if err := lu.repository.GetById(tx, loan); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to delete loans",
		}
	}
	// check if loan is returned before deleting
	if loan.LoanStatus_id != uint(entity.RETURNED_CODE) {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     errors.New("loan not returned yet"),
			Message: "failed to delete loan",
		}
	}
	if err := lu.repository.Delete(tx, loan); err != nil {
		return nil, &model.Response{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to delete loan",
		}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &model.Response{Code: http.StatusInternalServerError, Err: err, Message: "failed to add loan"}
	}
	return loan, &model.Response{Code: http.StatusOK, Message: "delete loan success"}
}

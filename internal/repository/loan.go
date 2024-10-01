package repository

import (
	"fmt"

	"github.com/thiccpan/library_information_system/internal/entity"
	"github.com/thiccpan/library_information_system/internal/model"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (u *LoanRepository) GetAll(db *gorm.DB, req *model.QueryLoanRequest, cond map[string]any) ([]entity.Loan, error) {
	tx := db.Joins("Book").Joins("User").Joins("LoanStatus")
	var loans []entity.Loan
	statusId, ok := cond["status_id"].(uint)
	if ok && statusId != 0 {
		fmt.Println(statusId)
		tx = tx.Where("loan_status_id = ?", statusId)
	}
	tx = tx.Find(&loans)
	if err := tx.Error; err != nil {
		return loans, err
	}
	return loans, nil
}
func (u *LoanRepository) GetById(db *gorm.DB, loan *entity.Loan) error {
	tx := db.Joins("User").Joins("Book").First(&loan)
	return tx.Error
}

func (u *LoanRepository) Add(db *gorm.DB, loan *entity.Loan) error {
	tx := db.Create(loan)
	return tx.Error
}

func (u *LoanRepository) UpdateById(db *gorm.DB, loan *entity.Loan) error {
	tx := db.Model(loan).Where("id = ?", loan.Id).Updates(loan)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (u *LoanRepository) Delete(db *gorm.DB, loan *entity.Loan) error {
	tx := db.Delete(loan)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

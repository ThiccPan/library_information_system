package repository

import (
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

func (u *LoanRepository) GetAll(db *gorm.DB, req *model.QueryLoanRequest) ([]entity.Loan, error) {
	var loans []entity.Loan
	tx := db.Joins("Book").Joins("User").Find(&loans)
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

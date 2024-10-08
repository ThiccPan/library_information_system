package repository

import (
	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (u *UserRepository) GetAll(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	tx := db.Omit("password").Find(&users)
	if err := tx.Error; err != nil {
		return users, err
	}
	return users, nil
}

func (u *UserRepository) GetByEmail(db *gorm.DB, user *entity.User) error {
	tx := db.First(&user, "email = ?", user.Email)
	return tx.Error
}

func (u *UserRepository) GetById(db *gorm.DB, user *entity.User, cond map[string]any) error {
	tx := db
	statusId, ok := cond["status_id"]
	if ok && statusId != 0 {
		tx = tx.Preload("Loans", u.Db.Where(&entity.Loan{LoanStatus_id: uint(statusId.(int))}))
	}
	if ok && statusId == 0 {
		tx = tx.Preload("Loans")
	}

	tx.First(&user)
	return tx.Error
}

func (u *UserRepository) Add(db *gorm.DB, user *entity.User) error {
	tx := db.Create(user)
	return tx.Error
}

func (u *UserRepository) UpdateById(db *gorm.DB, user *entity.User) error {
	tx := db.Model(user).Where("id = ?", user.Id).Updates(user)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (u *UserRepository) Delete(db *gorm.DB, user *entity.User) error {
	tx := db.Delete(user)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

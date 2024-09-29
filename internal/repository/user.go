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
	tx := db.Find(&user, "email = ?", user.Email)
	return tx.Error
}

func (u *UserRepository) Add(db *gorm.DB, user *entity.User) error {
	tx := db.Create(user)
	return tx.Error
}

func (u *UserRepository) UpdateById(db *gorm.DB, user *entity.User) error {
	tx := db.Model(user).Where("id = ?", user.Id).Updates(user)
	return tx.Error
}

func (u *UserRepository) Delete(db *gorm.DB, user *entity.User) error {
	tx := db.Delete(user)
	return tx.Error
}

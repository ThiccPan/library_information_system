package repository

import (
	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	Db *gorm.DB
}

func NewAuthorRepoImpl(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		Db: db,
	}
}

func (a *AuthorRepository) Add(db *gorm.DB, author *entity.Author) error {
	tx := db.Create(author)
	return tx.Error
}

func (a *AuthorRepository) GetAll(db *gorm.DB) ([]entity.Author, error) {
	var authors []entity.Author
	tx := db.Find(&authors)
	if err := tx.Error; err != nil {
		return authors, err
	}
	return authors, nil
}

func (a *AuthorRepository) GetById(db *gorm.DB, author *entity.Author) error {
	tx := db.First(&author)
	return tx.Error
}

func (a *AuthorRepository) UpdateById(db *gorm.DB, author *entity.Author) error {
	tx := db.Model(author).Where("id = ?", author.Id).Updates(author)
	return tx.Error
}

func (a *AuthorRepository) Delete(db *gorm.DB, author *entity.Author) error {
	tx := db.Delete(author)
	return tx.Error
}

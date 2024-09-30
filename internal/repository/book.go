package repository

import (
	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func NewBookRepoImpl(db *gorm.DB) *BookRepository {
	return &BookRepository{
		Db: db,
	}
}

func (u *BookRepository) GetAll(db *gorm.DB) ([]entity.Book, error) {
	var books []entity.Book
	tx := db.Joins("Author").Joins("Topic").Find(&books)
	if err := tx.Error; err != nil {
		return books, err
	}
	return books, nil
}
func (u *BookRepository) GetById(db *gorm.DB, book *entity.Book) error {
	tx := db.Joins("Author").Joins("Topic").First(&book)
	return tx.Error
}

func (u *BookRepository) Add(db *gorm.DB, book *entity.Book) error {
	tx := db.Create(book)
	return tx.Error
}

func (u *BookRepository) UpdateById(db *gorm.DB, book *entity.Book) error {
	tx := db.Model(book).Where("id = ?", book.Id).Updates(book)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (u *BookRepository) Delete(db *gorm.DB, book *entity.Book) error {
	tx := db.Delete(book)
	if tx.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

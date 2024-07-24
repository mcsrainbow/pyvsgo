package models

import "github.com/jinzhu/gorm"

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (repo *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	if err := repo.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (repo *BookRepository) CreateBook(book Book) error {
	err := repo.DB.Create(&book).Error
	return err
}

func (repo *BookRepository) GetBookByID(id uint) (Book, error) {
	var book Book
	err := repo.DB.First(&book, id).Error
	return book, err
}

func (repo *BookRepository) UpdateBook(book Book) error {
	err := repo.DB.Save(&book).Error
	return err
}

func (repo *BookRepository) DeleteBook(book Book) error {
	err := repo.DB.Delete(&book).Error
	return err
}

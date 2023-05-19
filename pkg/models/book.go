package models

import (
	"github.com/jinzhu/gorm"
	"github.com/suhel-kap/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var foundBook Book
	db := db.Where("ID=?", ID).Find(&foundBook)
	return &foundBook, db
}

func DeleteBook(ID int64) Book {
	var deleteBook Book
	db.Where("ID=?", ID).Delete(deleteBook)
	return deleteBook
}

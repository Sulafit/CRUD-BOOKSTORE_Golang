package models

import (
	"github.com/Sulafit/Go/SulaGO/project/db"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Description string `gorm:""json:"description"`
	Price       int64    `gorm:""json:"price"`
	Rating      int64    `gorm:""json:"rating"`
	Deleted gorm.DeletedAt
	

}

var DB *gorm.DB

func init() {

	DB = db.DbInit()
	DB.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {

	DB.Create(&b)
	return b
}
func GetAllBooks() []Book {
	var books []Book
	DB.Find(&books)
	return books
}
func GetAllByPriceDesc() []Book{
	var books []Book
	DB.Order("price DESC").Find(&books)
	return books
}

func GetAllByPriceAsc() []Book{
	var books []Book
	DB.Order("price ASC").Find(&books)
	return books
}

func GetBookById(id int64) (Book, *gorm.DB) {
	var book Book
	result := DB.Find(&book, id)
	// db := db.DB.First(&book, params["Id"])
	return book, result

}
func DeleteBook(id int64) (Book) {
	var book Book
	// DB.Where("ID=?", id).Delete(&book)
	DB.Delete(&book,id)
	// DB.Unscoped().Delete(&book, id)
	return book

}

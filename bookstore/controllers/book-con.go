package controllers

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"strconv"

	"github.com/Sulafit/Go/SulaGO/project/utils"
	"github.com/gorilla/mux"

	// "github.com/Sulafit/Go/SulaGO/project/db"

	"github.com/Sulafit/Go/SulaGO/project/models"
)

var NewBook models.Book

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsling")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(&books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println("Err")

}

func GetBooksDesc(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllByPriceDesc()
	res, _ := json.Marshal(&books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println("Err")

}

func GetBooksAsc(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllByPriceAsc()
	res, _ := json.Marshal(&books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println("Err")
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsling deletfunc")

	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(r, book)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsling")

	}
	// .Save(&ook{ID: 1, Name: "jinzhu", Age: 100})

	bookDetails, db := models.GetBookById(ID)
	
	if book.ID != 0 {
		bookDetails.ID = book.ID
	}
	if book.Name != "" {
		bookDetails.Name = book.Name
	}
	if book.Description != "" {
		bookDetails.Description = book.Description
	}
	if book.Price != 0 {
		bookDetails.Price = book.Price
	}
	if book.Rating != 0 {
		bookDetails.Rating = book.Rating
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

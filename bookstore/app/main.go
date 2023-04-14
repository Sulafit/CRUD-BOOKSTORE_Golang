package main

import (
	// "database/sql"
	"fmt"
	// "log"

	"net/http"
	// "strconv"
	"github.com/gorilla/mux"

	// "github.com/lib/pq"
	// "github.com/Sulafit/Go/SulaGO"
	"github.com/Sulafit/Go/SulaGO/project/controllers"
	// "github.com/Sulafit/Go/SulaGO/project/models"
	// "github.com/Sulafit/Go/SulaGO/project/models"

	// "github.com/Sulafit/Go/SulaGO/project/db"
	// "github.com/Sulafit/Go/SulaGO/project/models"

	_ "github.com/lib/pq"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to MyLibrary!") // replace this with your desired homepage content
// }

// ____________________________________________________________________________________________________________________________________
func main() {
	// db.DB.AutoMigrate(models.Books{})
	// m := &models.GetAllBooks()
	// fmt.Print(m)
	// models.init()

	r := mux.NewRouter()

	// r.HandleFunc("/", homeHandler).Methods()
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/desc", controllers.GetBooksDesc).Methods("GET")
	r.HandleFunc("/asc", controllers.GetBooksAsc).Methods("GET")

	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")

	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	// log.Fatal(http.ListenAndServe("localhost:8080", r))

	http.Handle("/", r)

	server := http.Server{
		Addr: ":8080",
	}

	fmt.Println("Server started listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

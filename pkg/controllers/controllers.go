package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irron21/book-management-system/pkg/models"
	"github.com/irron21/book-management-system/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookByID(ID)
	resp, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	resp, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	resp, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

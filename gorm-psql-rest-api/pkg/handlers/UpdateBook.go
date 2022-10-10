package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/mocks"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)
	updateBook.Id = id

	w.Header().Add("Content-Type", "application/json")
	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books[index] = updateBook
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mocks.Books[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/mocks"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	w.Header().Add("Content-Type", "application/json")
	// find and remove a book with specified id
	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mocks.Books)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/mocks"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find a book with this id and send it in response
	w.Header().Add("Content-Type", "application/json")
	for _, book := range mocks.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

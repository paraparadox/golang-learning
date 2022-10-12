package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"net/http"
	"strconv"
)

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find a book with this id and send it in response
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record not found"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

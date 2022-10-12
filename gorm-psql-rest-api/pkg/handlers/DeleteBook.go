package handlers

import (
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"net/http"
	"strconv"
)

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// find a book with specified id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record not found"))
		return
	}

	// delete that book
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

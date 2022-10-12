package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	updatedBook.Id = id

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record not found"))
		return
	}
	h.DB.Save(&updatedBook)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

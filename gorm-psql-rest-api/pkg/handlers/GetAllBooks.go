package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"net/http"
)

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

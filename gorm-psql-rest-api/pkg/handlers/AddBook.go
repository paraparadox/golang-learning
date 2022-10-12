package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"io"
	"log"
	"net/http"
)

func (h *handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// read the request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// append to book mocks
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

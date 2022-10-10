package handlers

import (
	"encoding/json"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/mocks"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	// read the request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// append to book mocks
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	// send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

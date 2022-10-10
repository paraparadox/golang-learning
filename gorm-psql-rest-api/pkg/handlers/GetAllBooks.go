package handlers

import (
	"encoding/json"
	"github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/mocks"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}

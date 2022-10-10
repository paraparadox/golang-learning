package mocks

import "github.com/paraparadox/golang-learning/gorm-psql-rest-api/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
	},
}

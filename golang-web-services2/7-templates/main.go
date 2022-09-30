package main

import (
	"bytes"
	"fmt"
	"golang-learning/golang-web-services2/7-templates/item"
	"golang-learning/golang-web-services2/7-templates/template"
	"net/http"
)

//go:generate hero -source=./template

var ExampleItems = []*item.Item{
	&item.Item{1, "rvasiliy", "Mail.ru Group"},
	&item.Item{2, "username", "freelancer"},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buffer := new(bytes.Buffer)
		template.Index(ExampleItems, buffer)
		w.Write(buffer.Bytes())
	})

	fmt.Println("starting server at :4000")
	http.ListenAndServe(":4000", nil)
}

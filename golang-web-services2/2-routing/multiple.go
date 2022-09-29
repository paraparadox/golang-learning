package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func RegularRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Request with average amount of logic\n")
}

func FastRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Request with high hitrate\n")
}

func ComplexRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Request with complex routing logic\n")
}

func main() {
	fastApiHandler := httprouter.New()
	fastApiHandler.GET("/fast/:id", FastRequest)

	complexApiHandler := mux.NewRouter()
	complexApiHandler.HandleFunc("/complex/", ComplexRequest).Headers("X-Requested-With", "XMLHttpRequest") // ajax

	stdApiHandler := http.NewServeMux()
	stdApiHandler.HandleFunc("/std/", RegularRequest)

	siteMux := http.NewServeMux()
	siteMux.Handle("/fast/", fastApiHandler)
	siteMux.Handle("/complex/", complexApiHandler)
	siteMux.Handle("/std/", stdApiHandler)

	fmt.Println("starting server at :4010")
	log.Fatal(http.ListenAndServe(":4010", siteMux))
}
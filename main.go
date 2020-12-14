package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Index :=")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func main() {
	handleRequests()
}

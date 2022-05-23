package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	ID    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "Camry", 1000},
	{2, "Nissan", "Pickup", 2000},
	{3, "VW", "Atlas", 20000},
	{4, "Honda", "Civic", 500},
}

func main() {
	fmt.Println("Booting up webserver")
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homehandler).Methods("GET")
	r.HandleFunc("/cars", carshandler).Methods("GET")
	err := http.ListenAndServe("localhost:8081", r)
	if err != nil {
		log.Fatal(err)
	}

}

func homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home page")
}

func carshandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	// Set up our routes
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/cars", returnCars).Methods("GET")
	r.HandleFunc("/cars/make/{make}", returnCarsByBrand).Methods("GET")
	r.HandleFunc("/cars/{id}", returnCarByID).Methods("GET")
	r.HandleFunc("/cars/{id}", updateCarByID).Methods("PUT")
	r.HandleFunc("/cars", createCarsHandler).Methods("POST")
	r.HandleFunc("/cars/{id}", deleteCarsHandler).Methods("DELETE")

	// Start the server and serve the routes
	err := http.ListenAndServe("localhost:8081", r)

	// If server craps out, hard crash and log
	if err != nil {
		log.Fatal(err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home page")
}

func returnCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func returnCarsByBrand(w http.ResponseWriter, r *http.Request) {
	// mux.Vars(r) comes as a map of key-value pairs
	vars := mux.Vars(r)
	carMake := vars["make"]
	cars := &[]Vehicle{}
	for _, car := range vehicles {
		if car.Make == carMake {
			*cars = append(*cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)

}

func returnCarByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert to string", err)
	}
	for _, car := range vehicles {
		if car.ID == carID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}

}

func updateCarByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert int to string", err)
	}
	var updatedCar Vehicle
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for k, v := range vehicles {
		if v.ID == carID {
			// Remove one vehicle
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
			// Add the new updated car
			vehicles = append(vehicles, updatedCar)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func createCarsHandler(w http.ResponseWriter, r *http.Request) {
	// Creating an empty newCar variable of type Vehicle
	var newCar Vehicle

	// Decoding the incoming body and assigning to newCar
	json.NewDecoder(r.Body).Decode(&newCar)

	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func deleteCarsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unabel to convert int to string", err)
	}
	for k, v := range vehicles {
		if v.ID == carID {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

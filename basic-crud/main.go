package main

/*
Create A Simple RESTful API With Golang
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json: "id, omitempty"`
	FirstName string   `json: "firstname, omitempty"`
	LastName  string   `json: "lastname, omitempty"`
	Address   *Address `json: "address, omitempty"`
}

type Address struct {
	City  string `json: "city, omitempty"`
	State string `json: "state, omitempty"`
}

var people []Person

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1])
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	//
	people = append(people, Person{ID: "1", FirstName: "Ben", LastName: "Huang", Address: &Address{City: "Cupertino", State: "CA"}})
	people = append(people, Person{ID: "2", FirstName: "Tom", LastName: "Liu", Address: &Address{City: "Newark", State: "CA"}})

	// Init router
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	// Http server
	log.Fatal(http.ListenAndServe(":8964", router))
}

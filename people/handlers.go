package people

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

// swagger:route GET /people Person GetPeople
// Returns a list of people
// responses:
//    200: peopleResponse
//    500: internalServerError
//    401: unauthorizedError
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !validateKey(w, r) {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

// swagger:route GET /people/{id} Person GetPerson
// Returns a Person
// responses:
//    200: personResponse
//    500: internalServerError
//    401: unauthorizedError
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !validateKey(w, r) {
		return
	}
	params := mux.Vars(r)
	for _, item := range people {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// swagger:route POST /people Person CreatePerson
// Create person from the database
// responses:
//    201: personCreatedResponse
//    500: internalServerError
//    401: unauthorizedError
// parameters:
// +name: person
//  in: body
//  description: A json that represents a person
//  required: true
//  type: Person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !validateKey(w, r) {
		return
	}
	w.WriteHeader(http.StatusCreated)
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

// swagger:route DELETE /people/{id} Person DeletePerson
// Delete person from the database
// responses:
//	202: acceptedResponse
//  500: internalServerError
//  401: unauthorizedError
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !validateKey(w, r) {
		return
	}
	w.WriteHeader(http.StatusAccepted)
	params := mux.Vars(r)
	for index, item := range people {
		if item.Id == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Person{})
}

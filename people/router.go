package people

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() {

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	router.Handle("/docs", getDocumentationHandler())
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":3333", router)
}

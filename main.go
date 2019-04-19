package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miun173/go-rest/handler"
	"github.com/miun173/go-rest/repo"
)

func init() {
	repo.InitDB()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", handler.GetPeople).Methods("GET")
	router.HandleFunc("/person", handler.GetPerson).Methods("GET")
	router.HandleFunc("/person", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handler.DeletePerson).Methods("DELET")
	router.HandleFunc("/person/{id}", handler.UpdatePerson).Methods("PUT")
	log.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

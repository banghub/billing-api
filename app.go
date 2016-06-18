package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/py150504/billingps/src/global"
	"github.com/py150504/billingps/src/people"
)

func main() {
	global.InitDB()
	people.InitPeople()
	log.Println("Run on : http://localhost:8080")

	r := mux.NewRouter()
	r.HandleFunc("/", global.Index)

	r.HandleFunc("/people", people.Read).Methods("GET")
	r.HandleFunc("/people/{id}", people.ReadDetail).Methods("GET")
	r.HandleFunc("/people", people.Create).Methods("POST")
	r.HandleFunc("/people/{id}", people.Delete).Methods("DELETE")

	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(global.NotFound)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}

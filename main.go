package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prasoonsoni/notes-backend-golang/controllers"
	"github.com/prasoonsoni/notes-backend-golang/db"
)

func main() {
	fmt.Println("Task Manager Backend GoLang")
	db.Connect()
	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Welcome to Task Manager Backend API</h1>"))
	})

	r.HandleFunc("/create", controllers.CreateTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
	http.Handle("/", r)

}

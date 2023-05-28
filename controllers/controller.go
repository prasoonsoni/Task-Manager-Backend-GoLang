package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prasoonsoni/notes-backend-golang/db"
	"github.com/prasoonsoni/notes-backend-golang/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Task")
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	// Decoding body and storing in variable
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(&models.Response{Success: false, Message: "Error Decoding Body"})
		return
	}

	// If any of the value is empty in body
	if task.IsEmpty() {
		json.NewEncoder(w).Encode(&models.Response{Success: false, Message: "Values Cannot Be Empty"})
		return
	}
	task.Completed = false

	result, err := db.TaskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(&models.Response{Success: false, Message: "Error Creating Task"})
		return
	}
	fmt.Println("Inserted ID:", result.InsertedID)
	json.NewEncoder(w).Encode(&models.Response{Success: true, Message: "Task Inserted Successfully"})
}

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prasoonsoni/notes-backend-golang/db"
	"github.com/prasoonsoni/notes-backend-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Task")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var id string = params["id"]
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	result, err := db.TaskCollection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Success: false, Message: "Error Deleting Task"})
	}
	fmt.Println("Inserted ID:", id)
	fmt.Println("Deleted Count:", result.DeletedCount)
	json.NewEncoder(w).Encode(&models.Response{Success: true, Message: "Task Deleted Successfully"})
}

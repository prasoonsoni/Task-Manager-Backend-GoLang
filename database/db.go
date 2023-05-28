package db

import (
	"context"
	"fmt"
	"log"

	"github.com/prasoonsoni/notes-backend-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://prasoonsoni:prasoon123@cluster0.zkyx4cg.mongodb.net/?retryWrites=true&w=majority"
const dbName = "task-manager-backend-golang"
const colName = "tasks"

var TaskCollection *mongo.Collection

// connection with MongoDB

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(uri)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Successful")

	TaskCollection = client.Database(dbName).Collection(colName)

	// Collection Instance
	fmt.Println("Collection instance is ready!!")
}

// Helper Functions - file

// Insert Task
func insertTask(task models.Task) {
	inserted, err := TaskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 task with ID:", inserted.InsertedID)
}

// Update Task
func updateTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set:": bson.M{"completed": true}}

	result, err := TaskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count:", result.ModifiedCount)
}

// Delete Task
func deleteTask(taskId string) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	result, err := TaskCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count: ", result.DeletedCount)
}

// Delete All Task
func deleteAllTask() {
	filter := bson.D{{}}
	deleteResult, err := TaskCollection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Tasks:", deleteResult.DeletedCount)
}
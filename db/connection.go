package db

import (
	"context"
	"fmt"
	"log"

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

func Connect() {
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

// Delete All Task
func deleteAllTask() {
	filter := bson.D{{}}
	deleteResult, err := TaskCollection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Tasks:", deleteResult.DeletedCount)
}

// Get All Tasks
func getAllTasks() []primitive.M {
	cursor, err := TaskCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var tasks []primitive.M

	for cursor.Next(context.Background()) {
		var task bson.M
		err := cursor.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	defer cursor.Close(context.Background())
	return tasks
}

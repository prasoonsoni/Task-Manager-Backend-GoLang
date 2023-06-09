package db

import (
	"context"
	"fmt"
	"log"

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

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

const (
	mongouri = "mongodb://hpcargo:27017"
)

func getconn() {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongouri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	// Rest of the code will go here
	getconn()
}

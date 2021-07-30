package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	mongouri = "mongodb://admin:secret@hpcargo:27017"
)

func getconn() {
	// Set client options
	// a := context.Background()
	// fmt.Println(a)
	// fmt.Println(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongouri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		client.Disconnect(ctx)
		fmt.Println("disconnected!")
	}()

	ash := Trainer{"Ash", 10, "Pallet Town"}
	collection := client.Database("test").Collection("trainers")

	insertResult, err := collection.InsertOne(ctx, ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	// Rest of the code will go here
	getconn()
}

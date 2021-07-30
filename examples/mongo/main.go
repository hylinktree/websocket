package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name                         string
	Age                          int
	City                         string
	VoltageR, VoltageS, VoltageT float64
}

const (
	mongouri = "mongodb://admin:secret@hpcargo:27017"
)

func randvoltage() float64 {
	const (
		min = 98.0
		max = 120.0
	)
	return min + rand.Float64()*(max-min)
}

func getconn() {
	const (
		timeout = 90
	)
	ticker := time.NewTicker(time.Second)
	go func() {
		i := 0
	p001:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(i, t)
				i++
			case <-time.After(time.Second * 3):
				fmt.Println("time is up")
				ticker.Stop()
				break p001
			}
		}
	}()
	// Set client options
	// a := context.Background()
	// fmt.Println(a)
	// fmt.Println(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
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

	for i := 0; i < 999999; i++ {
		ash := Trainer{"Ash", 10, "Pallet Town", randvoltage(), randvoltage(), randvoltage()}
		collection := client.Database("test").Collection("trainers")

		_, err := collection.InsertOne(ctx, ash)
		if err != nil {
			log.Fatal(err)
		}
		if (i % 1000) == 0 {
			fmt.Printf("%d records sent\n", i)
		}
	}

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

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

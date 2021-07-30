package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

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

// sample query
//  { $and: [ { voltager: { $gt: 119.999 } }, { voltager: { $exists: true } } ] }

func getconn() {
	// const (
	// 	timeout = 90
	// )
	// ticker := time.NewTicker(time.Second)
	// timer := time.After(time.Second * timeout)
	// go func() {
	// 	i := 0
	// p001:
	// 	for {
	// 		fmt.Println("looping")
	// 		select {
	// 		// case <-time.After(time.Second * 3): NOT WORK IF W/ TICKER
	// 		case <-timer:
	// 			fmt.Println("time is up", i)
	// 			ticker.Stop()
	// 			break p001
	// 		case t := <-ticker.C:
	// 			fmt.Println(i, t)
	// 			i++
	// 		}
	// 	}
	// }()
	// Set client options
	// a := context.Background()
	// fmt.Println(a)
	// fmt.Println(context.TODO())

	ctx := context.TODO()
	// As you set the timeout, it is expected the operation should be completed within the time!!
	// ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	// defer cancel()
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

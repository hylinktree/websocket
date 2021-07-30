package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	uri := "mongodb://admin:secret@" + *host + ":" + *port
	// "mongodb://admin:secret@hpcargo:27017"
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	fmt.Println("connecting to", uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		client.Disconnect(ctx)
		fmt.Println("disconnected!")
	}()
	var i = 0

	for ; i < *runs; i++ {
		// ash := Trainer{"Ash", 10, "Pallet Town", randvoltage(), randvoltage(), randvoltage()}
		collection := client.Database("test").Collection("trainers")

		_, err := collection.InsertOne(ctx, bson.D{
			{"family", "brian"},
		})

		// _, err := collection.InsertMany(ctx, []interface{}{
		// 	bson.D{
		// 		{"Name", "podcastResult.InsertedID"},
		// 	},
		// })
		if err != nil {
			log.Fatal(err)
		}
		if i != 0 && (i%1000) == 0 {
			fmt.Printf("%d records sent\n", i)
		}
	}
	fmt.Printf("total %d records sent\n", i)

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

var (
	host       = flag.String("host", "hpcargo", "host")
	port       = flag.String("port", "27017", "port")
	runs       = flag.Int("runs", 1, "total runs")
	database   = flag.String("db", "postgres", "the database for run")
	collection = flag.String("co", "meter101", "the collection for run")
)

func getconn2() {

	ctx := context.TODO()
	// As you set the timeout, it is expected the operation should be completed within the time!!
	// ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	// defer cancel()
	uri := "mongodb://admin:secret@" + *host + ":" + *port
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	fmt.Println("connecting to", uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		client.Disconnect(ctx)
		fmt.Println("disconnected!")
	}()
	var i = 0

	for ; i < *runs; i++ {
		// ash := Trainer{"Ash", 10, "Pallet Town", randvoltage(), randvoltage(), randvoltage()}
		collection := client.Database(*database).Collection(*collection)

		var err error

		_, err = collection.InsertOne(ctx, bson.D{
			{"voltageR", randvoltage()},
			{"voltageS", randvoltage()},
			{"voltageT", randvoltage()},
			{"timeStamp", time.Now().Unix()}, //makeJavaTimestamp(time.Now())},
		})

		if err != nil {
			log.Fatal(err)
		}
		_, err = collection.InsertMany(ctx, []interface{}{
			bson.D{
				{"voltageR", randvoltage()},
				{"voltageS", randvoltage()},
				{"voltageT", randvoltage()},
				{"timeStamp", time.Now().Unix()}}, //makeJavaTimestamp(time.Now())},
			bson.D{
				{"voltageR", randvoltage()},
				{"voltageS", randvoltage()},
				{"voltageT", randvoltage()},
				{"timeStamp", time.Now().Unix()}},
		})

		// _, err := collection.InsertMany(ctx, []interface{}{
		// 	bson.D{
		// 		{"Name", "podcastResult.InsertedID"},
		// 	},
		// })
		if err != nil {
			log.Fatal(err)
		}
		if i != 0 && (i%1000) == 0 {
			fmt.Printf("%d records sent\n", i)
		}
	}
	fmt.Printf("total %d records sent\n", i)

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func makeJavaTimestamp(t time.Time) int64 {
	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}
func main() {

	flag.Parse()
	getconn2()
}

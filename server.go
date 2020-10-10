package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

func main() {
	// Open Connection
	client, err := mongo.Connect(context.TODO(), "mongodb://$[username]:$[password]@$[hostlist]/$[database]?authSource=$[authSource]");

	if err != nil {
		log.Fatal(err)
	}

	// End Open Connection Code

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

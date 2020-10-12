package main

import (
	"context"
	"fmt"
	// "log"
	"time"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/chef-project"))

	defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
	}()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	// episodesCollection := database.Collection("episodes")

	podcast := Podcast{
    Title:  "The Polyglot Developer",
    Author: "Nic Raboy",
    Tags:   []string{"development", "programming", "coding"},
}

	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
		if err != nil {
	    panic(err)
	}
	fmt.Println(insertResult.InsertedID)
	http.ListenAndServe(":8080", nil)
}

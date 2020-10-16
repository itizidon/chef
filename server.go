package main

import (
	"context"
	"fmt"
	"log"
	"io/ioutil"
	"time"
	"net/http"
	// "encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty`
	Username string `bson:"username,omitempty"`
	Admin bool `bson:"admin,omitempty"`
	Firstname string `bson:"firstname,omitempty"`
	Lastname string `bson:"lastname,omitempty"`
	Age int `bson:"age,omitempty"`
	Email string `bson:"email,omitempty"`
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

	database := client.Database("chef-project")
	usersCollection := database.Collection("users")

	usersCollection.DeleteMany(ctx,bson.D{})
	user := User{
    Username:  "itizidon",
    Admin: true,
		Firstname: "Don",
		Lastname: "Ng",
		Age: 24,
		Email: "Don@email.com",
}

	insertResult, err :=usersCollection.InsertOne(ctx, user)

	r := mux.NewRouter()
	r.HandleFunc("/",handler).Methods("GET")

	r.HandleFunc("/lol",postHandler).Methods("POST")
	// http.HandleFunc("/", handler)
	fmt.Println(insertResult.InsertedID)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// jsn, err := ioutil.ReadAll(r.Body)
	if(err != nil){
		log.Fatal("error", err)
	}

	// json.Unmarshal(jsn, m)
	// fmt.Println(m)

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w)
}

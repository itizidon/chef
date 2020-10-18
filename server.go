package main

import (
	"context"
	"fmt"
	"log"
	"io/ioutil"
	"time"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	util "chef-project/util"
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

	r.HandleFunc("/newUser",newUserHandler).Methods("POST")
	r.HandleFunc("/createRecipe", newRecipe).Methods("POST")
	fmt.Println(insertResult.InsertedID)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func newRecipe(w http.ResponseWriter, r *http.Request){

}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
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
	jsn, err := ioutil.ReadAll(r.Body)
	if(err != nil){
		log.Fatal("error", err)
	}
	var data map[string]interface{}

	json.Unmarshal(jsn, &data)

	usersCollection.InsertOne(ctx,data)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w)
}

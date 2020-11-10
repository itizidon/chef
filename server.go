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

type JavaScript struct {
	Code  string
	Scope interface{}
}

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Admin bool `bson:"admin,omitempty"`
	Firstname string `bson:"firstname,omitempty"`
	Lastname string `bson:"lastname,omitempty"`
	Age int `bson:"age,omitempty"`
	Email string `bson:"email,omitempty"`
	Shop util.List
}

type RecipeQuery struct {
	RecipeKey string `bson:"recipekey,omitempty"`
	RecipeType string `bson:"recipetype,omitempty"`
}

type AllRecipes struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID string `bson:"userid,omitempty"`
	Recipename string `bson:"recipename,omitempty"`
	Time int `bson:"time,omitempty"`
	Ethnicity string `bson:"ethnicity,omitempty"`
	Method string `bson:"method,omitempty"`
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

	addRecipe := util.Recipe{
		"pasta",
		"mom's famous pasta",
		500,
	}

	newShop := util.List{}
	newShop.Insert(addRecipe)

	user := User{
    Username:  "itizidon",
    Admin: true,
		Firstname: "Don",
		Lastname: "Ng",
		Age: 24,
		Email: "Don@email.com",
		Shop: newShop,
}

	usersCollection.InsertOne(ctx, user)

	r := mux.NewRouter()
	r.HandleFunc("/",handler).Methods("GET")
	r.HandleFunc("/getRecipes", getRecipes).Methods("POST")
	r.HandleFunc("/newUser",newUserHandler).Methods("POST")
	r.HandleFunc("/createRecipe", newRecipe).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func newRecipe(w http.ResponseWriter, r *http.Request){
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
	allRecipes := database.Collection("generalRecipes")
	jsn, err := ioutil.ReadAll(r.Body)
	if(err != nil){
		log.Fatal("error", err)
	}

	var data AllRecipes

	json.Unmarshal(jsn, &data)

	allRecipes.InsertOne(ctx, data)
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
	var data User

	json.Unmarshal(jsn, &data)

	usersCollection.InsertOne(ctx,data)
}

func handler(w http.ResponseWriter, r *http.Request) {


}

func getRecipes(w http.ResponseWriter, r *http.Request){

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
	allRecipes := database.Collection("generalRecipes")
	jsn, err := ioutil.ReadAll(r.Body)
	if(err != nil){
		log.Fatal("error", err)
	}


	var data RecipeQuery
	json.Unmarshal(jsn, &data)

	returnedRecipes, err := allRecipes.Find(ctx,bson.M{data.RecipeKey: data.RecipeType})

	var allRecipesParsed []bson.M
	if err = returnedRecipes.All(ctx, &allRecipesParsed); err != nil {
    log.Fatal(err)
}


	// fmt.Println(allRecipes.Find(ctx, bson.M{}))
	// fmt.Println(allRecipesParsed)
	json.NewEncoder(w).Encode(allRecipesParsed)

	// fmt.Fprintf(w, allRecipesParsed)
}

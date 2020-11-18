package main

import (
	"context"
	"fmt"
	"log"
	// "io/ioutil"
	"time"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	gohandlers "github.com/gorilla/handlers"
	util "chef-project/util"

)

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
	generalRecipesCollection := database.Collection("generalRecipes")



	seedData := []AllRecipes{{UserID: "1", Recipename: "Pho", Time: 200, Ethnicity: "Viet", Method: "Broth"},{UserID: "2", Recipename: "Burger", Time: 200, Ethnicity: "American", Method: "BBQ"},{UserID: "3", Recipename: "Fish", Time: 500, Ethnicity: "American", Method: "Grill"}, {UserID:"4", Recipename: "Fries", Time: 200, Ethnicity: "American", Method: "Deep Fry"}, {UserID: "5",Recipename: "Sushi", Time: 300, Ethnicity: "Japanese"}, {UserID: "6", Recipename: "Wonton", Time: 400, Ethnicity:"Chinese"}}

	generalRecipesCollection.DeleteMany(ctx, bson.D{})
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
	for i := 0; i < len(seedData); i++ {
		seedRecipe:= AllRecipes{
			UserID: seedData[i].UserID,
			Recipename: seedData[i].Recipename,
			Time: seedData[i].Time,
			Ethnicity: seedData[i].Ethnicity,
			Method: seedData[i].Method,
		}
		generalRecipesCollection.InsertOne(ctx, seedRecipe)
	}

	r := mux.NewRouter()
	r.HandleFunc("/",handler).Methods("GET")
	r.HandleFunc("/getRecipes", getRecipes).Methods("POST")
	r.HandleFunc("/newUser",newUserHandler).Methods("POST")
	r.HandleFunc("/createRecipe", newRecipe).Methods("POST")
	r.HandleFunc("/getTags", getTags).Methods("GET")
	http.Handle("/", r)

	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:3000"}),
		gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(r)

	http.ListenAndServe(":8080", ch)
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

	var data AllRecipes

	json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
	var data User
	json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	var data *util.RecipeInfo

	json.NewDecoder(r.Body).Decode(&data)

	if data.RecipeKey == "get all" {
		returnedRecipes, err := allRecipes.Find(ctx,bson.M{})

		var allRecipesParsed []bson.M
		if err = returnedRecipes.All(ctx, &allRecipesParsed); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(allRecipesParsed)
	} else {

		searchRecipe := util.Queryify(data)
		fmt.Println(searchRecipe)
		returnedRecipes, err := allRecipes.Find(ctx,searchRecipe)

		fmt.Println(data.Recipename, data.Ethnicity)
		var allRecipesParsed []bson.M
		if err = returnedRecipes.All(ctx, &allRecipesParsed); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(allRecipesParsed)
	}
}

func getTags (w http.ResponseWriter, r *http.Request){
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

	// matchStage := bson.D{{"$match", bson.D{{"recipename",bson.D{{"$exists",true}}}}}}

	// groupStage := bson.D{{"$group", bson.D{{"recipename",bson.D{{"$exists",true}}, {"$total", bson.D{{"$addToSet", "$recipename"}}}}}}}

	// showInfoCursor, err := allRecipes.Aggregate(ctx, groupStage)

	// fmt.Println(showInfoCursor)
	returnedTags, err := allRecipes.Find(ctx,bson.M{})
	var allTagsParsed []bson.M
	if err = returnedTags.All(ctx, &allTagsParsed); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(allTagsParsed)
}

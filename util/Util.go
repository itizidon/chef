package util

import (
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

)

type RecipeInfo struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UserID []string `bson:"userid,omitempty"`
	Recipename []string `bson:"recipename,omitempty"`
	Time []int `bson:"time,omitempty"`
	Ethnicity []string `bson:"ethnicity,omitempty"`
	Method []string `bson:"method,omitempty"`
	RecipeKey string `bson:"recipekey,omitempty"`
}

type Recipe struct {
  Name string
  Description string
  Price int
}

type Node struct {
    Prev *Node
    Next *Node
    Key  Recipe
}

type List struct {
    Head *Node
    Tail *Node
}

func (L *List) Insert(Key Recipe) {
    list := &Node{
        Next: L.Head,
        Key:  Key,
    }
    if L.Head != nil {
        L.Head.Prev = list
    }
    L.Head = list

    l := L.Head
    for l.Next != nil {
        l = l.Next
    }
    L.Tail = l
}

func (l *List) Display() {
    list := l.Head
    for list != nil {
        fmt.Printf("%+v ->", list.Key)
        list = list.Next
    }
    fmt.Println()
}

func Display(list *Node) {
    for list != nil {
        fmt.Printf("%v ->", list.Key)
        list = list.Next
    }
    fmt.Println()
}

func ShowBackwards(list *Node) {
    for list != nil {
        fmt.Printf("%v <-", list.Key)
        list = list.Prev
    }
    fmt.Println()
}

func (l *List) Reverse() {
    curr := l.Head
    var Prev *Node
    l.Tail = l.Head

    for curr != nil {
        Next := curr.Next
        curr.Next = Prev
        Prev = curr
        curr = Next
    }
    l.Head = Prev
    Display(l.Head)
}

func appendToQuery(arr []string) bson.A{
    var data bson.A
    for _,v := range arr{
        data = append(data,v)
    }
    return data
}

func appendToQueryInt(arr []int) bson.A{
    var data bson.A
    for _,v := range arr{
        data = append(data,v)
    }
    return data
}
func Queryify(data *RecipeInfo) bson.M{

    var recipename = appendToQuery(data.Recipename)

    var ethnicity = appendToQuery(data.Ethnicity)

    var method = appendToQuery(data.Method)

    var time = appendToQueryInt(data.Time)

    result := bson.M{}

    if len(data.Recipename) != 0 {
        result["recipename"] = bson.M{"$in":recipename}
    }

    if len(data.Ethnicity) != 0 {
        result["ethnicity"] = bson.M{"$in": ethnicity}
    }

    if len(data.Method) != 0 {
        result["method"] = bson.M{"$in": method}
    }

    if len(data.Time) != 0 {
        result["time"] = bson.M{"$in": time}
    }
    return result
}

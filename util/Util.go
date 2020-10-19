package util

import (
	"fmt"
)

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

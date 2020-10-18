package util

import (
	"fmt"
)

type Recipe struct {
  name string
  description string
  price int
}

type Node struct {
    prev *Node
    next *Node
    key  Recipe
}

type List struct {
    head *Node
    tail *Node
}

func (L *List) Insert(key Recipe) {
    list := &Node{
        next: L.head,
        key:  key,
    }
    if L.head != nil {
        L.head.prev = list
    }
    L.head = list

    l := L.head
    for l.next != nil {
        l = l.next
    }
    L.tail = l
}

func (l *List) Display() {
    list := l.head
    for list != nil {
        fmt.Printf("%+v ->", list.key)
        list = list.next
    }
    fmt.Println()
}

func Display(list *Node) {
    for list != nil {
        fmt.Printf("%v ->", list.key)
        list = list.next
    }
    fmt.Println()
}

func ShowBackwards(list *Node) {
    for list != nil {
        fmt.Printf("%v <-", list.key)
        list = list.prev
    }
    fmt.Println()
}

func (l *List) Reverse() {
    curr := l.head
    var prev *Node
    l.tail = l.head

    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
    Display(l.head)
}

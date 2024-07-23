package main

import "fmt"

// Simple linked list implementation
type linkedList struct {
    item string
    next *linkedList
}

func main() {
    head := linkedList{item: "Head", next: nil}
    tail := linkedList{item: "Tail", next: nil}

    head.next = &tail
    fmt.Println(head.next)

    for i := head; i.next != nil; i = *i.next {
        fmt.Printf("Value at i is %v\n", i.item)
    }  
}

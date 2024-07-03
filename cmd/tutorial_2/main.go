package main

import "fmt"

func main() {
    var myNum string = "072214124";
    fmt.Println(myNum);

    greeting("Donald Duck")
}

// Its good that in Go you don't have to place function definitions before you
// call them
func greeting(name string) {
    fmt.Printf("\nHello %v", name);
}

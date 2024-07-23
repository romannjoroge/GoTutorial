package main

import "fmt"

type student struct {
    name string
    regno string
}

func (s student) registerForClass(class string) {
    fmt.Printf("%v has registered for %v succesfully\n", s.regno, class)
}

func (s student) greet(name string) {
    fmt.Printf("Hello %v\n", name)
}

type person interface {
    greet(name string) 
}

func greetSomeone(p person, name string) {
    p.greet(name)
}

func main() {
    var student1 = student{name:"Alice Doe", regno: "P16/2000/2089"}
    student1.registerForClass("Operating Systems")

    greetSomeone(student1, "Omollo Okech")
}

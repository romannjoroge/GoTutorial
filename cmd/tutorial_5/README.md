## Structs and Interfaces

The time has finally come to start looking at the big boy stuff in Go no more simple shit!

Structs are similar to structs in C as I have seen so far. They allow you to combine data that is related to each other. A classic example is that every person has a name and an age. Instead of having 2 seperate variables to store data of a person or 2 seperate lists to store the data of multiple people create a struct that combines the name and age into one struct and use a variable of that type to represent a person or a list of values of that type.

The syntax of defining a struct is as follows:

```go
type person struct { //structure is type <struct_name> struct
    name string // fields have structure of name type
    age uint8
}

// Syntax for creating a variable of struct
var p1 person // Creates a struct that will have default values in this case name is empty string since that dfault of string and age 0

// Create person and give it values
var p1 = person{name: "John", age: 0}
```

To access the field of a struct we use a . e.g `p1.name`.

There are also anonymous structs where you define and use it at the same place. An example would be:

```go
var p2 = struct{
    name string
    age uint8
}{"John", 12}
```

It turns out that structs

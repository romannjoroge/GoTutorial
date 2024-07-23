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

It turns out that structs aren't really the same to their C counterparts. They can have methods attached to them making them behave more like a class. Adding methods to a struct has abit of wonkie syntax ngl. What we do is that we assign the function to the struct when creating it. An example of a method for the person struct would be as follows:

```go
func (p person) greeting() string {
    return "Hello " + p.name
}

// General syntax
func <name-of-instance struct-name> <func-name> (args) return-type {
    // Body
}
```

## Interfaces

So we know the need for interfeces which is explained really well in the Pragmatic Programmer book. Interfaces allow a way for someone to define that multiple types can perform the same actions or have similar properties. It is a good alternative to inheritance where you specify only what is shared between the types and don't have to inherit everything.

To create an interface we do the following:

```go
type <interface-name> interface {
    <method-interface-name>(args) return-type
}

// Let's say that we want a greeting interface
type greeting interface {
    greet() string
}

// We know define that types of this interface must have a greet method that returns a string
```

Looks like you don't need to do something like extending an interface for a struct to belong to it. As long as the struct implements the method then it belongs to the interface which know that I think about it more it makes sense

## Review

Structs and interfaces in Go aren't bad. The way methods are added to structs was abit wierd to me but one can argue that it avoids cases where class definitions become so massive because you have to define the method in the class. Looking at it from this view can make this syntax make alot of sense.

Interfaces are good and I like that you don't have to specify that a struct implements an interface for it to be considered as following the rules of the interface.



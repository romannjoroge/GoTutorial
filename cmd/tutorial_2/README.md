## const, var and types in GO

So like in many other languages Go has constant variables and variables whose values can change later in the code. I assume you use const to denote constant variables and var for variables whose value can change. For each of these variables there are types i.e string etc.

I don't really see a need to go much more in depth than this unless some wierd GO shit happens!

Maybe something worth noting is that there is no float type you have to specify if its either float64 or float32.

The format for creating variables is

```go
var <varname> = value
var <varname> <type>

const <varname> = value // For constants you have to give an inital value

<varname> := value // a shorthand where you don't need to write var all the time

var <varname1>, <varname2> <type> = <value1>, <value2> // you can also define multiple variables in one line and this can have the different variations shown above
```

This might be similar to C but you can do arithmetic operations between values of different types i.e you can't add an int to a float and stuff like that. If you do need to do this though you can type cast which is done by using the name of the type as a function which is similar to how Python does it I think. An example is:

```go
var num1 int = 10
var float1 float32 = 2.0

var div = float1 / float32(num1)
```

To create multiline strings use the ` character and in a similar way to C strings are created using ".

You can get the lenght of a string using the `len()` function which again is similar to python. Okay scratch that Go has started some nonesense that honestly reminds me of some shit C would pull. So len() counts the number of bytes in a string. This is all good if you're only using utf-8 strings but as soon as you introduce some fancy unicode characters you're fucked cause unicode characters can be stored in 2 bytes. So the len of something like π could be 2 or something greater.

To actually count the number of characters in a string (or runes as they are referred to in Go [I've heard this before from somewhere]) You need to import the `unicode/utf-8` package and use the `utf8.RuneCountInString()` function.

Runes are represented using ' character in a similar way to C. Printing a rune prints its code e.g printing 'a' prints 97 which is similar behaviour to assembly of all things.

Something interesting about format strings is that we use the %v to represent a variable to place in the format string. We use %v for values of all kinds of types. To print a format string we use `fmt.Printf`. A caveat of this is that it behaves in a similar way to printf in C where it does not put a \n automatically at the end of the string!

## A little runt

Nothing big so far but I hope this pattern of behaving in a similar way to C or assembly stops. Was expecting Go to have a little more of an abstarction than this. Anyway its not too bad but caught me off guard abit.

## Functions and control statements

So let's define the full format of something that is in every language.....functions! So the func keyword is used to define a function and the syntax is as follows:

```go
func <func-name> () {
    // function body
}
```

Well functions need arguements so to add those we do the following:

```go
func <func_name> (<varname> type, <varname2> type, ...) {
    // Function body
}
```

Functions in Go can also return multiple things. Functions in most languages can also do this but it seems especially important in Go since people return errors rather than throwing them! So in a function that can encounter an error you would typically see in the list of return types the error type. An example

```go
func <funcname> (args) (return-type1, return_type2 ...) {
    // Function body

}
```

So if such a function encounters something that would in other languages throw an error you instead return it and if there is no error you return a non initialized variable of type error or nil. An example of this is found in the code.

To create an error first import the `errors` package and run the `errors.New(<description_of_err>)` function. To view the error messages of an error object run the `error.Error()` function.


### If statements

So in a similar fashion to python the conditions in if statements are not wrapped in (). An example of an if condition would be:

```go
if condition1 {
    // Body
} else if condition2 {
    // Body
} else {
    // Condition
}
```

### Switch statement

Abit of an odd one is switch. I don't think I've seen this syntax before but it is definetly interesting and could save alot of boilerplate. One can use a switch without giving it a value to look at i.e

```go
switch {
    case cond1:
        // Do stuff
    case cond2:
        // Do stuff
    default:
        // Do stuff
}
```

One interesting thing about this is that it allows the conditions to be looking at different variables for example cond1 could be checking if x < 3, cond2 could be checking if y > 2 and so on. It acts exactly the same way as if-else if-else. Also the break at the end of each case is optional resulting in much cleaner code.

On the other hand you can use the case like normal and give it a variable and the cases to look at possible values of the variable i.e

```go
switch <variable> {
    case <value1>:
        // Do Stuff
    case <value2>, <value3>: // This means execute if variable is either value2 or value 3
        // Do stuff
    default:
        // Do stuff
}
```

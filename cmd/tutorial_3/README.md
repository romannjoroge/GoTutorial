## Arrays, Maps, Loops

So I was originally planning on getting my hands abit dirty with Go before coming to this part of the series. I had a feeling that knowing about Arrays and looping would be important but I decided to see how far I could go without them. I was able to print some text in a box and saw that to print a paragraph in a box it would be really helpful to use an array and loops so here I am!

Fuck my worst nightmare! The arrays are of a fixed length! I was hoping they'd be like vectors or growable arrays but nah this is different! I will continue watching but may decide on a workaround.

So similar to many languages array are 0 indexed and are indexed using []. In a similar way to python you can provide a start and end value to the [] i.e `array[1:3]` gets the values of the array from index 1 to index 2. The final index is not included. The location in memory of where a member in the array , already reminds me of C sadly, can be indexed using the &. This gives you a pointer.

To create an array we do the following:

```go
var <arrayName> [array-size]arrayElementType // Just initialize but no initial value
var myArr [3]int

var <arrayName> [array-size]arrayElemType = [array-size]arrayElemType{elements...} // Give initial value
var myArray2 [3]int = [3]int{1, 2, 3}

var := [arraySize]arrayElemType{elements....} 
```

Thank God Go has an implementation of growable arrays called Slices. Slices add the functionality of being able to append values to it. An example of how to create a slice

```go
var <slice-name> []ArrayElemType = []{elements...}

// To append to a slice use the append function
mySlice := []int{1, 2, 3} // When you use := there is no need to use var keyword
mySlice = append(mySlice, 4)
```

The `len()` function gives the length of an array while the `cap()` function gives you the number of elements the array can hold i.e if you have an array [5]int{1, 2} could have a length of 2 and a capacity of 5

### Maps

Okay lets speedrun this. To create a map you can do one of the following:

```go
var myMap = map[type-of-key]type-of-value{key: value, ....}

var myMap map[keyType]valueType = make(map[keyType]valueType)
```

To access a value of the map do the following `map[key]`. One odd thing about maps in Go is that they always return a value so if you do map[keyThatDoesNotExist] you will get the default value of the value type i.e 0 for a map with int as value. To check if the map actually has that key it returns a second value that is true if element exists and false if it doesn't i.e

```go
var val, ok = map[key]
if ok {
    // Value exists in map
} else {
    // value does not exist
}
```
To remove an element from a map you can use the `delete` function.

### Loops

You can loop over a map or an array using a for loop. The syntax of for loop is:

```go
for var:= range array-or-map {
    // Do something
}
```

When looping through a map giving one value to the for loop gets the keys and adding a second value lets you get the values i.e: (In the case of an array the first value is index while second is value.
```go
for key := map {
    // Stuff
}

for key, value := range map {
    // Stuff
}

for i, value := range array {
    // Stuff
}
```

Damn Go doesn't have a while loop wtf. So to achieve a while we still use for in a similar way we would do a while loop i.e:

```go
for condition { // Breaks when condition becomes false
    // Stuff
}

// Infinite loop
for { 
    // Stuff
}
```

Typical for loop that we all love and are used to:

```go
for i:=0; condition with i; i++ {
    // Stuff
}
```

The examples from this tutorial can be found at the mini_project_1 folder!

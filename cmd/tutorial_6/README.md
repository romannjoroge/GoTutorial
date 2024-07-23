## Pointers

Pointers are variables that hold locations in memory. They are similar to phone numbers which store something that you can use to reach a friend. To declare a pointer use the following syntax:

```go
var pointer *int32;
var pointer *int32 = new(int32)

// General syntax
var <pointer-name> *<type>
```

The new command works in the same way as malloc. It gives you a free memory address i.e memory that isn't storing anything that is big enough to store a value of the indicated type.

To get the value stored in the memory location being pointed at we use the * operator. To get the memory location of a variable we use the & operator

```go
i := 10
p := &i
fmt.Println(*p)
```



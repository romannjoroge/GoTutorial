## Generics

Generics allow you to define that a function or a struct can be used on values of many types. In Go when using generics you have to define the types that the generic could possible be i.e this function can accept either a int or a float32. It sounds abit similar to interfaces with this respect. An example of using generics is as follows:

```go
func sumSlice[T int | float32](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return v
}

type car [T int | float32] struct {
    name string
    weight T
}
```

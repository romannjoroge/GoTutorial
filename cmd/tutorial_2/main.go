package main

import "fmt"
import "errors"

func main() {
    var myNum string = "072214124";
    fmt.Println(myNum);

    greeting("Donald Duck")

    // Dividing
    num1 := 4
    num2 := 2
    var div, error  = div2Nums(num1, num2)
    if error != nil {
        fmt.Printf("Got a %v error\n", error.Error())
    } else {
        fmt.Printf("Function divided sucessfully and got %v as a result\n", div);
    }
}

// Its good that in Go you don't have to place function definitions before you
// call them
func greeting(name string) {
    fmt.Printf("\nHello %v", name);
}


func div2Nums(num1 int, num2 int) (int, error) {
    var error error;
    if num2 == 0 {
        error = errors.New("Cannot Divide by Zero")
        return 0, error
    }

    return num1 / num2, error
}

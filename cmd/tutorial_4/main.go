package main

import (
    "fmt"
    "strings"
)

func main() {
    classStrings := []string{
        "John Doe ",
        "A1523342025 ",
        "Bsc. Computer Science ",
    }
    
    var strBuilder strings.Builder
    for _, str:= range classStrings {
      strBuilder.WriteString(str)
    }

    var regString = strBuilder.String()
    fmt.Printf("Student Details Are => %v\n", regString)
}

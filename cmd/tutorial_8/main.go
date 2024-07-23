package main

import (
	"fmt"
	"time"
)

func main() {
    var c = make(chan int)
    go getData(c)
    process(c)
}

func getData(c chan int) {
    for i:=0; i < 10; i++ {
        addData(c, i)
    }
    close(c)
}

func addData(c chan int, data int) {
   c <- data
   time.Sleep(time.Millisecond * 500)
}

func process(c chan int) {
    for i:= range(c) {
        fmt.Println(i)
    }
}

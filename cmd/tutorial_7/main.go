package main

import (
	"fmt"
	"time"
    "sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var results []int32

func dbCall(i int32){
    time.Sleep(time.Millisecond * 500)
    fmt.Printf("%v succesfully written\n", i)
    m.Lock()
    results = append(results, i)
    m.Unlock()
    wg.Done()
}

func main() {
    for i := int32(0); i < 10; i++ {
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
    fmt.Println(results)
}

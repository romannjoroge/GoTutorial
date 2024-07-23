## Concurrency

Something interesting to note is that concurrency and parallel execution are NOT the same. Parallel execution takes place when your CPU has multiple cores so one core can be dealing with one task and at the same time another core is dealing with another task.

Concurrency on the other hand also allows for a situation where one core executes task 1 for some time then jumps to start executing task 2. Something similar to pipelining where as task 1 is being executed task 2 could be being read in memory.

Most of the times concurrency in Go is achieved through parallel execution but its not the case all the times.

## Goroutines

We use Goroutines to achieve concurrency in Go. When using Go routines you need to consider concepts like race conditions and mutual exclusion. But before we look at that let's see how to add a function to a go routine. To do that add go infront of the function call and add it to a wait group

```go
let w = sync.WaitGroup{}
wg.Add(1)
go function()
```

What we've done above is that we've added the function to a wait group. Wait groups contain function that can be run in parallel. If we want to use the results or the side effects of running the functions of a wait group we need to wait for them to finish running. We do this using the ~wg.Wait()~ function

I guess in summary wait groups allow you to define which set of functions you want to be able to run concurrently. One issue of concurrency is that the functions might try editing the same resource at the same time or might try reading from a resource as it is being written. To solve these issues we need to ad mutual exclusion to the code. To do this we can do the following:

```go
var m = sync.RWMutex{}

m.Lock()
// Edit resource
m.Unlock()

m.RLock()
// read resource
m.RUnlock()
```

The ~Lock()~ method locks the resource. This makes other functions that need to use the resource to first have to wait until the resource is unlocked to then use it. The ~RLock()~ method indicates that a function is currenltly reading a resource. If the resource is being read by another function at the same time it doesn't block but if the resource is currently being written i.e it is ~Lock()~ it stops a new function from going to read it. At the same time ~Lock()~ waits until all ~RLock()~ are ~RUnlock()~ before allowing the write.

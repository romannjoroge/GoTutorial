## Channels

While working with go routines I realised that I couldn't do something like this:

```go
var i = go func()
```

I had thought that goroutines are similar to async await in JS and that I could simply "await" for the result of a go routine. You cannot do this and I instead ended up creating a global slice that the goroutine would return its values to then I would process the result. What I did is similar to how channels work.

Instead of having an array that stores the output of the goroutine I could store this output in a channel. Once I push the result to a channel I could have something listen to data being pushed to the channel and process the data.

The benefit of this is that unlike in the array approach where I had to wait for all of the go routines to complete before I could read the array I can know read and work on data as it comes! The following is the syntax:

```go
var channel = make(chan <type-of-data-in-channel>) // Channels can only hold data of one type
channel <- 1 // Syntax for pushing a value to a channel
close(channel) // Closing a channel indicates that your routine is done adding values to it
var i = <- c // syntax for getting vaue from channel
func(i) 
```

If your process is listening from multiple channels and you want to do something different depending on what channel the data came from you can use the ~select~ statement.


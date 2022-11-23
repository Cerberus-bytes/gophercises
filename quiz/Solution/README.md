# Notes

## Arguments
```golang
// Declare the flag/switches
csvFileName := flag.String("csv", "problems.csv", "A CSV file in the format  of 'question,answer'")
timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
```

The [flag](https://pkg.go.dev/flag) module allows you to pass arguments from the command line into the golang program. The format is:
```
flag.[Type](name, default_values, help_message)
```

After initializing flags, you will need to parse all your options:
```golang
flag.Parse()
```

The arguments are saved as pointers and can be accessed as such.
```golang
// Create file handle
file, err := os.Open(*csvFileName)
if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to open CSV file %s\n", *csvFileName)
    os.Exit(1)
}
```

## Timers
```golang
timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
<-timer.C
```
The [time](https://pkg.go.dev/time) module has a sub-module, [timer](https://pkg.go.dev/time#Timer). Timer represents a single event. When the Timer expires, the current time will be sent on C. Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.

[In depth explanation for timers](https://gobyexample.com/timers)

## Goroutines and Channels
```golang
answerCh := make(chan string)
go func() {
    var answer string
    fmt.Scanf("%s", &answer)
    answerCh <- answer // Send user input to answer channel
}() // Put `()` at the end to call this immediatly
```
A goroutine is a lightweight thread managed by the Go runtime. [More Information on Concurrency](https://go.dev/tour/concurrency/1)
Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`. For example:
```golang
ch <- v     // Send v to channel ch
v := <-ch   // Receive from ch, and assign value to v
```
Like maps and slices, channels must be created before use:
```golang
ch := make(chan int)
```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or conditional values.
The `select` statement lets goroutines wait on multile communication operations. A `select` blocks until one of the cases can run, then executes that case.
```golang
select {
case <-timer.C: // If you get message from timer channel
    fmt.Printf("\nYou scored %d / %d\n", correct, len(problems))
    return
case answer := <- answerCh: // If you get an answer from the answer channel
    if answer == p.Solution {
        correct++
    }
}
```
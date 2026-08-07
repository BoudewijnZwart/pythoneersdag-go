# Concurrency & parallelism assignments

The Go language was designed with concurrency and parallelism in mind from the beginning. These assignments will teach you the basics.

## Assignment 1 - Normal synchronous code

To understand the advantage of concurrency we start with a small program that run synchronously. Open de files 01_sync/main.go, shared/utils/colors.go and shared/utils/toys.go and try to understand what this code will do. You can run the code from the folder this markdown file is in by executing:

```bash
go run ./01_sync/
```

The code runs fairly slow. Is there a better way and where can we make the biggest improvements?

## Assignment 2 - Concurrency without parallelism

Open the file 02/concurrent/main.go. For this assignment we will not use parallelism just yet. The line

```go
runtime.GOMAXPROCS(1)
```

set the amount of OS threads that will do work to one, disabeling paralellism for now. Try to make the creation of toys run concurrently. Use goroutines and wait groups.


## Assignment 3 - Parrallelism with Mutex

If a Go program has multiple cores available the goroutines will be executed not only concurrently but also in parallel, if possible. You don't need to do anything extra to enable this, the Go runtime will do this for you.
Running code in parallel comes with some extra responsibility for the developer. You have to make sure that all the goroutines do not try to use the same memory at the same time. This can cause data races.

Go has a tool to detect data races. With the following command you can see the possible data race we have created in assignment 3:

```bash
    go run -race /03_parallel_mutex
```

Where is the data race happening?

Try to solve this problem by using a Mutex. See https://go.dev/tour/concurrency/9 for more information on Mutexes in Go.

## Assignment 4 - Parallelism with Channels

A more common way of solving memory access isues in Go is by using channels. Channels transport data between goroutines. The goroutine must specify if it reads from or rights to a channel. Channels are designed in a way that makes sure that even if two goroutines are trying to write to them or read from them at the exact same time, a ordering is still enforced.

See https://go.dev/tour/concurrency/2 for more information on channels.




# When to Use Concurrency

- The goroutine is the core concept in Go’s concurrency model

> How Programme work normally.
    - Process
    - Thread

- Single process can have multiple thread. Process will have it's own memory, That will not
Share with other Process. But Threads will share memory in side Process.

Process-A
    - thread-01
    - thread-02
    - thread-03
    ...

Process-B
    - thread-04
    - thread-05
    - thread-06
    ...
------------------------
> Goroutines are lightweight processes managed by the Go runtime.
> When a Go program starts, the Go runtime creates a number of threads and launches a single goroutine to run your program.

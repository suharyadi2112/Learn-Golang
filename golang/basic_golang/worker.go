package main

import (
    "fmt"
    "time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
    for task := range tasks {
        fmt.Printf("Worker %d started task %d\n", id, task)
        time.Sleep(time.Second) // Simulate a time-consuming task
        results <- task * 2      // Send the result back to the main goroutine
        fmt.Printf("Worker %d finished task %d\n", id, task)
    }
}

func main() {
    numWorkers := 3
    numTasks := 10

    // Create channels for tasks and results
    tasks := make(chan int, numTasks)
    results := make(chan int, numTasks)

    // Launch a pool of worker goroutines
    for i := 1; i <= numWorkers; i++ {
        go worker(i, tasks, results)
    }

    // Generate tasks and send them to the task channel
    for i := 1; i <= numTasks; i++ {
        tasks <- i
    }

    // Wait for all results to be received from the workers
    for i := 1; i <= numTasks; i++ {
        result := <-results
        fmt.Printf("Received result %d\n", result)
    }

    // Close the channels to signal the workers to exit
    close(tasks)
    close(results)
}
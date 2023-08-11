package main

import (
    "fmt"
    "runtime"
)

func main() {
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)

    fmt.Printf("Initial memory stats: %+v\n", mem)

    // Allocate a bunch of memory
    var data [][]byte
    for i := 0; i < 100000; i++ {
        d := make([]byte, 10000)
        data = append(data, d)
    }

    // Force garbage collection
    runtime.GC()

    // Get memory stats again
    runtime.ReadMemStats(&mem)
    fmt.Printf("Final memory stats: %+v\n", mem)
}

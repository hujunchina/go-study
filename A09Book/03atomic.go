// This sample program demonstrates how to use the atomic
// package to provide safe access to numeric types.
package main

import (
"fmt"
"runtime"
"sync"
"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines.
	counter2 int64

	// wg is used to wait for the program to finish.
	wg2 sync.WaitGroup
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg2.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg2.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter2(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg2.Done()

	for count2 := 0; count2 < 2; count2++ {
		// Safely Add One To Counter.
		atomic.AddInt64(&counter2, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}

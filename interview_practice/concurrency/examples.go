package concurency

import (
	"fmt"
	"math/rand"
	"sync"
)

// ConcurentForLoop runs For loop in concurrent way
func ConcurentForLoop(size int) {
	var wg sync.WaitGroup
	wg.Add(size) // tell the 'wg' WaitGroup how many threads/goroutines that are about to run

	arr := make([]byte, size)
	rand.Read(arr) // Create an array of rnadom bytes that will be process.

	fmt.Println("Running for loop ...")
	for i := 0; i < size; i++ {
		// Spawn a thread for each iteration of loop.
		// Pass 'i' into the goroutine's function
		//   in order to make sure each goroutine uses
		//   uses different independent value for 'i'.
		go func(i int) {
			// At the end of goroutine tell the 'wg' WaitGroup that another thread is completed.
			defer wg.Done()

			v := arr[i]
			fmt.Printf("i: %v, val: %v\n", i, v)
		}(i)
	}
	fmt.Println("Doing other stuf...") // This part of the code will run before wg.Wait() pause the call stack.

	// Wait for 'wg.Done()' to be executed the number of times specified in the wg.Add() call (pause the call stack).
	// 'wg.Done()' should be executed exact same times as specified in wg.Add().
	// If the numbers do not match, 'wg.Wait()' will either hang infinitely or throw a panic error.
	wg.Wait()

	fmt.Println("Execution is finished")
}

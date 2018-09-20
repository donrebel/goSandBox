package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
)

// ================
// Squaring numbers
// ================
// Consider a pipeline with three stages.
// The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
// The gen function starts a goroutine that sends the integers on the channel and closes the channel when all the values
// have been sent:
func gen(nums []byte) <-chan byte {
	out := make(chan byte)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// The second stage, sq, receives integers from a channel and returns a channel that emits the square of each received integer.
// After the inbound channel is closed and this stage has sent all the values downstream, it closes the outbound channel:
func sc(in <-chan byte) <-chan byte {
	out := make(chan byte)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func genVals(size int) []byte {
	vals := make([]byte, size)
	rand.Read(vals)
	return vals
}

func squaringNumbers(size int) {
	c := gen(genVals(size))
	for x := range sc(c) {
		fmt.Printf("%v ", x)
	}
	fmt.Println()
}

// ==============
// Fan IN Fan OUT
// ==============
// Multiple functions can read from the same channel until that channel is closed; this is called fan-out.
// This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.
// A function can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that's closed when all the inputs are closed. This is called fan-in.
// We can change our pipeline to run two instances of sq, each reading from the same input channel.
// We introduce a new function, merge, to fan in the results:
func fanInFanOut(size int) {
	in := gen(genVals(size))
	c1 := sc(in)
	c2 := sc(in)

	for x := range merge(c1, c2) {
		// fmt.Printf("%v ", x)
		fmt.Println(x)
	}
	fmt.Println()
}

// The merge function converts a list of channels to a single channel by starting a goroutine for each inbound channel
// that copies the values to the sole outbound channel. Once all the output goroutines have been started, merge starts
// one more goroutine to close the outbound channel after all sends on that channel are done.
// Sends on a closed channel panic, so it's important to ensure all sends are done before calling close.
// The sync.WaitGroup type provides a simple way to arrange this synchronization:
func merge(cs ...<-chan byte) <-chan byte {
	var wg sync.WaitGroup
	out := make(chan byte)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan byte, cn int) {
		for x := range c {
			out <- x
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for i, c := range cs {
		fmt.Println(len(c))
		go output(c, i)
	}
	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// GoRoutines are not garbage collected, they must be exit by they own.
// So if downstream resources are not able to read So, in order to avoid resource leaks
//

// ExecPattern used to execute patterns for testing purposes in main()
func ExecPattern() {
	fanInFanOut(10)
}

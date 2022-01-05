package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			// Defer a call to wg.Done() to indicate that the background goroutine has
			// completed when this function returns. Behind the scenes this decrements
			// the WaitGroup counter by 1 and is the same as writing wg.Add(-1)
			defer wg.Done()
			fmt.Printf("hello from a goroutine %i\n", i)
		}(i)
	}

	// Wait() blocks until the WaitGroup counter is zero --- essentially blocking until all
	// goroutines have completed.
	wg.Wait()

	fmt.Println("all finished\n")
}

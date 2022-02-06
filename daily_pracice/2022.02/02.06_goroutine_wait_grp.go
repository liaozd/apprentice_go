package main

import (
	"fmt"
	"sync"
)

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			//defer mut.Unlock()
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("counter = %d", counter)
}

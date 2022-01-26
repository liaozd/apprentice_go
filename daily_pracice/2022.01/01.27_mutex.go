package main

import (
	"fmt"
	"sync"
	"time"
)

//https://shockerli.net/post/golang-pkg-mutex/
func main() {
	var mutex sync.Mutex

	wait := sync.WaitGroup{}

	fmt.Println("Locked")
	mutex.Lock()

	for i := 1; i <= 3; i++ {
		wait.Add(1)

		go func(i int) {
			fmt.Println("Not lock:", i)

			mutex.Lock()
			fmt.Println("Locked:", i)

			time.Sleep(time.Second)

			mutex.Unlock()
			fmt.Println("Unlocked:", i)

			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	mutex.Unlock()
	fmt.Println("Unlocked")

	wait.Wait()
}

/*
Locked
Not lock: 3
Not lock: 1
Not lock: 2
Unlocked
Locked: 3
Unlocked: 3
Locked: 1
Unlocked: 1
Locked: 2
Unlocked: 2
*/

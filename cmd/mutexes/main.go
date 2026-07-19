package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mutex sync.Mutex
	count int
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for range 10000 {
		wg.Go(func() {
			// LockからUnlockまでが排他
			counter.mutex.Lock()
			defer counter.mutex.Unlock()
			counter.count++
		})
	}

	wg.Wait()

	fmt.Printf("counter.count: %d\n", counter.count)
}

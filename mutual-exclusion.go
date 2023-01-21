package main

import (
	"fmt"
	"sync"
)

var sharedVar int
var mu sync.Mutex

func increment() {
	mu.Lock()
	sharedVar++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sharedVar)
}

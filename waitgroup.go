package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers done")
}

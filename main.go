package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// spin up workers
	numWorkers := 50
	for i := 0; i < numWorkers; i++ {

		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}

func worker(id int) {
	fmt.Println(id, " worker starting job")
	time.Sleep(time.Second)
	fmt.Println(id, " worker done")
}

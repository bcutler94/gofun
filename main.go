package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 50
	jobs := make(chan int, numJobs)
	results := make(chan bool, numJobs)

	// spin up workers
	numWorkers := 1
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- bool) {
	for range jobs {
		fmt.Println(id, " worker starting job")
		time.Sleep(time.Second)
		fmt.Println(id, " worker done")
		results <- true
	}
}

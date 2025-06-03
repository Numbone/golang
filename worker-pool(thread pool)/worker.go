package worker_pool_thread_pool_

import (
	"fmt"
	"time"
)

func worker_pool() {
	t := time.Now()
	const jobsCount, workerCount = 15, 15
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("result #%d :value = %d\n", i+1, <-results)
	}
	fmt.Printf("time elapsed: %s\n", time.Since(t))
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		//fmt.Println("worker", id, "started job", j)
		//fmt.Println("worker", id, "finished job", j)
		results <- j * j
	}
}

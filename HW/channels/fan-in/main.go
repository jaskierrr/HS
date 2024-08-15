package main

import (
	"fmt"
	"sync"
)

const (
	numJobs    = 6
	numWorkers = 3
)
//fan-in
func main() {

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// заполняю канал подзадач
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		//wg.Add(1) ??? имеет ли смысл
		go func() {
			// каждый воркер берет подзадачу из канала
			for job := range jobs {
				results <- job * job
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// обработка результатов
	for result := range results {
		fmt.Println(result)
	}
}

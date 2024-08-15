package main

import (
	"fmt"
	"sync"
	"time"
)

type Job func()

func worker(wg *sync.WaitGroup, id int, jobs <-chan Job) {
	defer wg.Done()
	// воркеры куртятся пока не выполнят все джобы
	for job := range jobs {
		fmt.Printf("worker %d starting job\n", id)
		job()
		fmt.Printf("worker %d finish job\n", id)
	}
}

func main() {
	jobsCount := 7
	jobs := make(chan Job, jobsCount)
	var wg sync.WaitGroup

	// Добавляю джобы в очередь
	go func() {
		for j := 1; j <= jobsCount; j++ {
			job := func() {
				time.Sleep(2 * time.Second)
				fmt.Printf("job %d is done\n", j)
			}
			jobs <- job
		}
		close(jobs)
	}()

	// Запускаю воркеров
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i, jobs)
	}

	// Закрытие канала означает, что больше джоб не появится
	wg.Wait()
	fmt.Println("All jobs have been finish")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	C chan struct{}
}

// добавляем горутину в семафор
func (sem *Semaphore) Acquire() {
	sem.C <- struct{}{}
}

// выпускаем горутину из семафора
func (sem *Semaphore) Release() {
	<-sem.C
}

const maxGoroutines = 7

func main() {
	// максимальное количество одновременно выполняющихся горутин

	sem := Semaphore{
		C: make(chan struct{}, maxGoroutines),
	}

	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			sem.Acquire()
			defer sem.Release()

			fmt.Printf("running task: %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)
const maxGo int = 3

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

func main() {
	wg := sync.WaitGroup{}
	sem := Semaphore{
		C: make(chan struct{}, maxGo),
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		
		go func (i int)  {
			sem.Acquire()

			defer wg.Done()
			defer sem.Release()

			printNumber(i)
		}(i)

	}
	wg.Wait()
}

func printNumber(n int) {
  time.Sleep(time.Second)
  fmt.Println(n)
}

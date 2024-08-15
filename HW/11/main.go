package main

import (
	"fmt"
)

// Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.
// FAN IN
func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)
	// реализовать вот здесь

	go func() {
		done := make(chan struct{})

		for _, item := range chs {
			go func(item <-chan int) {
				for itemIn := range item {
					mergedCh <- itemIn
				}

				done <- struct{}{}
				
			}(item)
		}

		for range chs {
			<-done
		}

		close(mergedCh)
	}()

	return mergedCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}

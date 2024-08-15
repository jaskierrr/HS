package main

import (
	"context"
	"fmt"
	"time"
)


func main() {
	ctx := context.Background()
	doneCh := make(chan bool)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer close(doneCh)
	defer cancel()

	go ticker(ctx, doneCh)

	doneFunc(doneCh)

}

func ticker(ctx context.Context, doneCh <-chan bool) {
	//ctx, cancel := context.WithTimeout(parentCtx, time.Second*2)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		case <-doneCh:
			fmt.Println("done chan")
			return
		default:
			fmt.Println("tick")
			time.Sleep(time.Second * 1)
		}
	}
}


func doneFunc(doneCh chan bool) {
	time.Sleep(4 * time.Second)
	doneCh <- true
}

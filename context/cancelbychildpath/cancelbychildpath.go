package main

import (
	"context"
	"fmt"
	"time"
)

// 取消子go中的子go
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go subgoFirst(ctx)

	time.Sleep(5 * time.Second)

	fmt.Println("notify exit")
	cancel() // 不发生阻塞

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")
	}
}

func subgoFirst(parent context.Context) {
	ctx, _ := context.WithCancel(parent)
	go subgoSecond(ctx)

	i := 1
	for {
		select {
		case <-parent.Done():
			fmt.Printf("subgoFirst exit:%s\n", parent.Err())
			return
		default:
			fmt.Printf("looping subgoFirst times %d\n", i)
			time.Sleep(1 * time.Second)
		}
		i++
	}
}

func subgoSecond(parent context.Context) {
	i := 1
	for {
		select {
		case <-parent.Done():
			fmt.Printf("subgoSecond exit:%s\n", parent.Err())
			return
		default:
			fmt.Printf("looping subgoSecond times %d\n", i)
			time.Sleep(1 * time.Second)
		}
		i++
	}
}

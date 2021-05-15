package main

import (
	"context"
	"fmt"
	"time"
)

// 取消子go中的子go
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("pass ctx from main to subgoFirst")
	go subgoFirst(ctx)

	time.Sleep(5 * time.Second)

	cancel()
	<-ctx.Done()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")
	}
}

func subgoFirst(parent context.Context) {
	ctx, _ := context.WithCancel(parent)
	fmt.Println("pass ctx from subgoFirst to subgoSecond")
	go subgoSecond(ctx)

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("looping subgoFirst")

		select {
		case <-parent.Done():
			fmt.Printf("subgoFirst exit:%s\n", parent.Err())
			return
		default:
		}
	}
}

func subgoSecond(parent context.Context) {

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("looping subgoSecond")

		select {
		case <-parent.Done():
			fmt.Printf("subgoSecond exit:%s\n", parent.Err())
			return
		default:
		}
	}
}

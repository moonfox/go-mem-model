package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	_ = cancelMutiGo
	_ = cancelMutiGoWithDiffCtx
	cancelMutiGo()
}

// 用同一个父ctx，取消拥有相同 ctx 的 多个goroutine
func cancelMutiGo() {
	ctxParent, cancelParent := context.WithCancel(context.Background())
	ctxChild, cancelChild := context.WithCancel(ctxParent)

	for i := 0; i < 5; i++ {
		go cancelGo(ctxChild, i, cancelChild)
	}

	time.Sleep(5 * time.Second)
	cancelParent()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")
	}
}

// 用同一个父ctx，取消拥有不同 ctx 的 多个goroutine
func cancelMutiGoWithDiffCtx() {
	ctxParent, cancelParent := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		ctxChild, cancelChild := context.WithCancel(ctxParent)
		go cancelGo(ctxChild, i, cancelChild)
	}

	time.Sleep(5 * time.Second)
	cancelParent()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")
	}
}

func cancelGo(ctx context.Context, num int, cancel context.CancelFunc) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("the number of goroutine: %d [%d]\n", num, i)

		select {
		case <-ctx.Done():
			fmt.Printf("%d canceled goroutine: [%d]: Why? %s \n", num, i, ctx.Err())

			return
		default:
		}
		i++
	}
}

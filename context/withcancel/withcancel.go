package main

import (
	"context"
	"fmt"
	"time"
)

func sleepRandom_1(stopChan chan struct{}) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 1: %d\n", i)

		i++
		if i == 5 {
			fmt.Println("cancel sleep random 1")
			stopChan <- struct{}{}
			break
		}
	}
}

func sleepRandom_2(ctx context.Context) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 2: %d\n", i)
		i++

		select {
		case <-ctx.Done():
			fmt.Printf("Why? %s\n", ctx.Err())
			fmt.Println("cancel sleep random 2")
			return
		default:
		}
	}
}

func main() {

	ctxParent, cancelParent := context.WithCancel(context.Background())
	ctxChild, cancelChild := context.WithCancel(ctxParent)

	stopChan := make(chan struct{})

	go sleepRandom_1(stopChan)
	go sleepRandom_2(ctxChild)

	select {
	case <-stopChan:
		fmt.Println("stopChan received")
	}
	cancelParent()
	cancelChild() // 可不写，写是为了不显示语法提示
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Continue...")
	}
}

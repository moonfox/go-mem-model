package main

import (
	"fmt"
	"sync"
	"time"
)

/* 是否可以读取没有解锁的资源
没有解锁的临界区的共享资源可以被没锁的goroutine读取
*/
func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	var ch = make(chan struct{}, 1)
	wg.Add(2)
	a := 0
	go func() {
		lock.Lock()
		a = 1
		ch <- struct{}{}
		// 不能通过sleep来控制语句的执行顺序，在-race中不起作用
		time.Sleep(1 * time.Second)
		lock.Unlock()
		fmt.Println("lock.Unlock()")
		wg.Done()
	}()

	go func() {
		// 1. 加入通道是为了确保 a = 1 赋值完成后 ，再进行打印
		// 2. a的值在解锁之前被读取

		fmt.Println(<-ch)
		fmt.Println(a) // => 1
		// a = 2
		wg.Done()
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

/* 是否可以读取没有解锁的资源
没有解锁的临界区的共享资源可以被没锁的goroutine读取
*/
func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan struct{})
	a := 0
	wg.Add(2)

	go func() { // Go1
		lock.Lock()
		a = 1
		ch <- struct{}{} // 确保Go2对a的读取，是在Go1对a写入之后
		// 不能通过sleep来控制语句的执行顺序，在-race中不起作用
		// time.Sleep(1 * time.Second)
		<-ch
		lock.Unlock()
		fmt.Println("lock.Unlock()")
		wg.Done()
	}()

	// fmt.Println(a) 的读取，在事件中的顺序是
	// 1.给a赋值{a=1} -> 2.读取a{Println(a)} -> 3.解锁{lock.Unlock()}
	// 通过加入对 通道的写入与读取，使并发事件按指定顺序执行

	go func() { // Go 2
		// 1. 加入通道是为了确保 a = 1 赋值完成后 ，再进行打印
		// 2. a的值在解锁之前被读取
		// a = 2 //在这里不是并发安全的

		fmt.Println(<-ch)
		// 假设 在未解锁时，不能对 a 进行读取，那么 Go2被阻塞，无法向 ch 写入数据
		// 那么 Go1 就会一直阻塞在 	<-ch，此时形成了 deadlock，程序异常。
		// 实际情况是程序可以正常执行，也就证明了可以在未解锁时，读取a
		fmt.Println(a)   // => 1
		ch <- struct{}{} // 确保Go1对lock.Unlock()的解锁，是在Go2对a的读取之后
		// a = 2 //在这里是并发安全的
		wg.Done()
	}()

	wg.Wait()
}

// Package main provides ...
package main

import (
	"fmt"
	"sync"
)

type Once struct {
	done uint32
	m    sync.RWMutex
}

// 同一把锁，不能在没解锁的情况下再次上锁，特别是在同一个goroutine中，这样做会锁死自己。
// 其它goroutine也无法为其解锁(很难)。这种给已经上锁的锁再次上锁的情况是很容易发生的
// 且不易被发现，特别是在一个函数调用别一个函数的时候
func (o *Once) Do(f func()) {

	o.m.RLock()
	defer o.m.RUnlock()
	if o.done == 0 {
		// 调用 doSlow(f) 时，因为还没有对  o.done 解锁(需要等到Do返回)
		// 所以 doSlow()无法进行加锁，陷入等待，等待 o.done进行解锁，所以陷入了死循环
		// 导致整个进程僵死
		// fatal error: all goroutines are asleep - deadlock!
		o.doSlow(f)
	}

}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer func() { o.done = 1 }()
		f()
	}
}

func main() {
	var once Once
	var wg sync.WaitGroup
	wg.Add(1)
	j := 2

	go once.Do(func() {
		fmt.Println("hello")
		fmt.Println(j)
		wg.Done()
	})

	wg.Wait()
}

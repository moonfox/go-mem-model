// Package main provides ...
package main

import (
	"fmt"
	"sync"
	"time"
)

// Once is an object that will perform exactly one action.
//
// A Once must not be copied after first use.
type Once struct {
	done uint32
	m    sync.Mutex
	m2   sync.Mutex
}

// 同一把锁，不能在没解锁的情况下再次上锁，特别是在同一个goroutine中，这样做会锁死自己。
// 其它goroutine也无法为其解锁(很难)。这种给已经上锁的锁再次上锁的情况是很容易发生的
// 且不易被发现，特别是在一个函数调用别一个函数的时候
// 所以在读取时使用另外一个锁进行读取
func (o *Once) Do(f func()) {
	// if atomic.LoadUint32(&o.done) == 0 {
	// 	o.doSlow(f)
	// }

	// o.m.RLock()
	// defer o.m.RUnlock()
	// if o.done == 0 {
	// 	o.doSlow(f)
	// }

	o.m2.Lock()
	defer o.m2.Unlock()
	if o.done == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		// defer atomic.StoreUint32(&o.done, 1)
		defer func() { o.done = 1 }()
		f()
	}
}

func main() {
	var once Once
	j := 2

	go once.Do(func() {
		fmt.Println("hello")
		fmt.Println(j)
	})

	go once.Do(func() {
		fmt.Println(j)
	})
	time.Sleep(1 * time.Second)

}

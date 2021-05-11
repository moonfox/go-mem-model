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
}

func (o *Once) Do(f func()) {
	// if atomic.LoadUint32(&o.done) == 0 {
	// 	// Outlined slow-path to allow inlining of the fast-path.
	// 	o.doSlow(f)
	// }

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
		fmt.Println(j)
	})

	go once.Do(func() {
		fmt.Println(j)
	})
	time.Sleep(2 * time.Second)

}

package main

import "sync"

var a, b int
var lock = sync.RWMutex{}

func f() {
	lock.Lock()
	a = 1
	b = 2
	lock.Unlock()
}

func g() {
	// lock.RLock()
	print(b)
	print(a)
	// lock.RUnlock()
}

func main() {
	go f()
	g()
}

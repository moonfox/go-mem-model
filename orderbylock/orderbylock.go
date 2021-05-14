package main

import (
	"sync"
)

// 对于任何 sync.Mutex 或 sync.RWMutex 类型的变量 l 以及 n < m ，
// 对 l.Unlock() 的第 n 次调用在对 l.Lock() 的第 m 次调用返回前发生
// 即：对于第二次锁定，要先解锁才行

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

// 可保证打印出 "hello, world"。该程序首先（在 f 中）对 l.Unlock() 进行第一次调用，
// 然后（在 main 中）对 l.Lock() 进行第二次调用，最后执行 print 函数
func main() {
	l.Lock()
	go f()
	l.Lock() // 发生阻塞，直到 f() 中执行l.Unlock()
	print(a)
}

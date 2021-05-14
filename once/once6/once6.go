package main

import "sync"

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

// 调用 twoprint 会打印两次 "hello, world"
// 第一次对 twoprint 的调用会运行一次 setup
// 会打印两次是重点，思考为什么可以打印两次
func twoprint() {
	go doprint()
	go doprint()
}

// 不加race不一定打印两次
// 没准一次都不打印
func main() {
	twoprint()
}

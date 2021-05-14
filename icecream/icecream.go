package main

import (
	"fmt"
)

type IceCreamMaker interface {
	// Great a customer.
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
	if b.name != "Ben" {
		panic("my name is Ben")
	}
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
	if j.name != "Jerry" {
		panic("my name is Jerry")
	}
}

// 输出名的名字不正确的情况
// Jerry says, "Hello my name is Ben"
// Ben says, "Hello my name is Jerry"
// 正确的应该是
// Ben says, "Hello my name is Ben"
// Jerry says, "Hello my name is Jerry"

// 这是因为我们在 maker = jerry 这种赋值操作的时候并不是原子的，
// 只有对 single machine word 进行赋值的时候才是原子的，虽然这个看上去只有一行，
// 但是 interface 在 go 中其实是一个结构体，它包含了 type 和 data 两个部分，
// 所以它的复制也不是原子的，会出现问题
func main() {
	var ben = &Ben{name: "Ben"}
	var jerry = &Jerry{"Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		maker.Hello()
	}
}

package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

// 这里不保证在 main 中对 done 的写入的监测， 蕴含对 a 的写入也进行监测，
// 因此该程序也可能会打印出一个空字符串
// 即 先写入 done，再写入 a，就会造成 main 读取到 done 时，a 可能还没有被写入

// 更糟的是，由于在两个线程之间没有同步事件
// 因此无法保证对 done 的写入总能被 main 监测到，即还没有同步到内存
// main 中的循环不保证一定能结束
func main() {
	go setup()
	for !done {
	}
	print(a)
}

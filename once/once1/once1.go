package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var icons map[string]string
var loadIconsOnce MyOnce

type MyOnce struct {
	done uint32
	m    sync.Mutex
}

func (o *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *MyOnce) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		// defer func() { o.done = 1 }()
		f()
	}
}

func main() {
	// go Icon("spades.png")
	// go Icon("diamonds.png")
	go func() { fmt.Println(Icon("spades.png")) }()
	go func() { fmt.Println(Icon("diamonds.png")) }()
	// go fmt.Println(Icon("diamonds.png"))
	// go loadIconsOnce.Do(loadIcons)
	// go loadIconsOnce.Do(loadIcons)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(Icon("clubs.png"))
}

func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcons() {
	icons = make(map[string]string)
	icons["spades.png"] = "spades.png"
	icons["hearts.png"] = "hearts.png"
	icons["diamonds.png"] = "diamonds.png"
	icons["clubs.png"] = "clubs.png"
	fmt.Println("loading finished")
}

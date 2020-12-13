package main

import (
	"sync"
)

// 对于sync.Mutex来说 n<m， 调用n次Unlock要先行发生于m次Lock
var l sync.Mutex
var a string

func f() {
	a = "hello world"
	l.Unlock()
}
func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}

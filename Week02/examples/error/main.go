package main

import (
	"fmt"
	"time"
)

func main() {
	Go(hello)
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func hello() {
	fmt.Println("hello")
	panic("panic hello")
}

// Go 开启goruntine
func Go(x func()) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	x()
}

package main

import (
	"sync/atomic"
)

func main() {
	var i int32
	atomic.AddInt32(&i, 3)

}

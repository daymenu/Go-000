package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wait sync.WaitGroup

var counter int32 = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		wait.Add(1)
		go myRoutine(routine)
	}
	wait.Wait()
	fmt.Printf("Final Counter : %d \n ", counter)
}

func myRoutine(id int) {
	for count := 0; count < 2; count++ {
		atomic.AddInt32(&counter, 1)
	}
	wait.Done()
}

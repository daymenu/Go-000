package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

var counter int = 0

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
		value := counter
		value++
		time.Sleep(1 * time.Nanosecond)
		counter = value
	}
	wait.Done()
}

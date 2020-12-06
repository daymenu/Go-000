package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var counter int

func main() {
	for r := 1; r <= 2; r++ {
		wait.Add(1)

		go routine(r)

	}

	wait.Wait()
	fmt.Printf("Final Counter : %d\n", counter)
}

func routine(id int) {
	for count := 0; count < 2; count++ {
		counter = counter + 1
	}
	wait.Done()
}

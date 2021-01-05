package main

import (
	"fmt"
	"time"

	"github.clm/daymenu/Go-000/Week06/sliding"
)

func main() {

	var opts []sliding.CounterOption
	opts = append(opts, sliding.WindowTime(4*time.Second))
	c := sliding.NewCounter(opts...)

	c.Inc(500)
	time.Sleep(200 * time.Millisecond)
	c.Inc(500)
	time.Sleep(200 * time.Millisecond)

	fmt.Println(c.Sum(time.Now()))
	fmt.Println(c.Avg(time.Now()))
}

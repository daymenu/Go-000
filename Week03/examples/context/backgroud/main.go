package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("ctx.Done: ", ctx.Done())
	fmt.Println("ctx.Err: ", ctx.Err())
	fmt.Println(ctx.Deadline())
}

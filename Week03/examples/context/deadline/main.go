package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 2 * time.Second

func main() {

	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	cancel()
}


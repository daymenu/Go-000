package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 20 * time.Millisecond

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)

	defer cancel()
	select {
	case <-time.After(1 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("ctx.Done()", ctx.Err())
	}
}

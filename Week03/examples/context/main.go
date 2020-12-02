package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	go sellMoive(ctx, "一路向西")

	select {
	case <-ctx.Done():
		fmt.Println("取消了")
		cancelFunc()
	}
}

func sellMoive(ctx context.Context, moiveName string) {
	fmt.Println("start sell moive")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("sell: ", moiveName)
			time.Sleep(time.Microsecond * 50)
		}
	}
}

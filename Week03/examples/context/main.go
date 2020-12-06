package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	name := "name"
	ctx = context.WithValue(ctx, name, "lili")
	go sellMoive(ctx, "一路向西")

	select {
	case <-ctx.Done():
		fmt.Println("取消了")
		cancelFunc()
	}
}

func sellMoive(ctx context.Context, moiveName string) {
	name := ctx.Value("name")
	fmt.Println("start sell moive, name:", name)
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

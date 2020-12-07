package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var errg errgroup.Group

// 基于 errgroup 实现多个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println(ctx, cancel)
}

func shoutdown(quit <-chan os.Signal) {
	switch <-quit {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("退出")
	default:
		fmt.Println("其他信号")
	}
}

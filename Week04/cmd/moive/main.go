package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/daymenu/Go-000/Week04/internal/moive"
	"golang.org/x/sync/errgroup"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	egroup, ctx := errgroup.WithContext(context.Background())

	// 启动影片服务
	egroup.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("moive quit")
				return ctx.Err()
			default:
				moive.App(appCtx)
			}
		}
	})

	// 启动信号监听
	egroup.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			cancel()
			return fmt.Errorf("用户退出")
		case <-ctx.Done():

		}
		fmt.Println("信号中断退出")
		return ctx.Err()
	})
	if err := egroup.Wait(); err != nil {
		log.Fatalln(err)
	}
}

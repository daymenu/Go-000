package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/daymenu/Go-000/Week03/work/ad"
	"github.com/daymenu/Go-000/Week03/work/moive"
	"golang.org/x/sync/errgroup"
)

// 基于 errgroup 实现多个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
// 该服务是一个影片服务
// 主要有两大服务 影片服务及广告服务
func main() {

	moiveErrGroup, _ := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(context.Background())
	// 启动广告服务
	moiveErrGroup.Go(func() error {
		err := ad.AppServe(ctx)
		cancel()
		fmt.Println("ad quit")
		return err
	})
	// 启动影片服务
	moiveErrGroup.Go(func() error {
		err := moive.AppServe(ctx)
		cancel()
		fmt.Println("moive quit")
		return err
	})

	// 监听退出信号
	moiveErrGroup.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			// 调用应用程序 context 取消所有程序
			cancel()
			fmt.Println("signal quit")
			return nil
		case <-ctx.Done():
			return nil
		}
	})

	if err := moiveErrGroup.Wait(); err != nil {
		cancel()
		fmt.Println("影片站点出错了:", err)
	}

	fmt.Println("影片站点所有服务已经退出了....")

}

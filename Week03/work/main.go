package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daymenu/Go-000/Week03/work/ad"
	"github.com/daymenu/Go-000/Week03/work/moive"
	"golang.org/x/sync/errgroup"
)

// 基于 errgroup 实现多个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
// 该服务是一个影片服务
// 主要有两大服务 影片服务及广告服务
func main() {
	startErrGroup, _ := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(context.Background())

	// 启动影片服务
	moiveServer := moive.Server{
		Server: &http.Server{
			Addr: ":9090",
		},
	}
	startErrGroup.Go(func() error {
		return moiveServer.Serve(ctx)
	})
	log.Println("启动影片服务")

	// 启动广告服务
	adServer := ad.Server{
		Server: &http.Server{
			Addr: ":9091",
		},
	}
	startErrGroup.Go(func() error {
		return adServer.Serve(ctx)
	})
	log.Println("启动广告服务")
	// 等待所有服务启动
	if err := startErrGroup.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("所有服务已经启动")

	// 处理退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()

	log.Println("接受到退出信号")

	quitCtx, quitCancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer quitCancel()

	stopErrGroup, _ := errgroup.WithContext(context.Background())
	stopErrGroup.Go(func() error {
		return adServer.Shutdown(quitCtx)
	})

	stopErrGroup.Go(func() error {
		return moiveServer.Shutdown(quitCtx)
	})

	if err := stopErrGroup.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("所有服务已经退出")
}

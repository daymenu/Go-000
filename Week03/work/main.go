package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var serverErrGroup errgroup.Group

// 基于 errgroup 实现多个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	serverErrGroup.Go(func() error {
		return gatewayServer(ctx)
	})

	serverErrGroup.Go(func() error {
		return appServer(ctx)
	})

	if err := serverErrGroup.Wait(); err != nil {
		log.Fatal(err)
	}

	defer cancel()
	fmt.Println("server is start")

	select {
	case <-quit:
		fmt.Println("")
	}
}

// 一个app服务
func appServer(ctx context.Context) error {
	srv := &http.Server{
		Addr:    ":80",
		Handler: nil,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	// err := fmt.Errorf("app")
	// return fmt.Errorf("server error: %w", err)
	return nil
}

// 一个网关服务
func gatewayServer(ctx context.Context) error {
	// err := fmt.Errorf("gateway")
	// return fmt.Errorf("server error: %w", err)
	return nil
}

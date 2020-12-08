package ad

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/daymenu/Go-000/Week03/work/code"
)

type adErr struct {
	err  error
	code code.ErrCode
}

func (err adErr) Error() string {
	return fmt.Sprintf("code: %d, err: %v", err.code, err.err)
}

func (err adErr) UnWrap() error {
	return err.err
}

// Server 影片服务
type Server struct {
	Server *http.Server
}

// Serve 启动服务
func (m *Server) Serve(ctx context.Context) error {
	addr := m.Server.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return &adErr{err: err, code: code.ServerErrCode}
	}
	go m.Server.Serve(ln)
	return nil
}

func (m *Server) startLog() error {
	err := fmt.Errorf("日志启动错误")
	return &adErr{err: err, code: code.LogErrCode}
}

// Shutdown 服务退出
func (m *Server) Shutdown(ctx context.Context) error {
	// 执行自己的退出操作
	fmt.Println("ad server quit")
	// 执行server的退出
	return m.Server.Shutdown(ctx)
}

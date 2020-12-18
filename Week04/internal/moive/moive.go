package moive

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/daymenu/Go-000/Week04/api/errcode"
)

// Service service
type Service struct {
	root map[string]http.HandlerFunc
}

// AddRoute 添加路由
func (s *Service) AddRoute(uri string, h http.HandlerFunc) {
	s.root[uri] = h
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	if h, ok := s.root[uri]; ok {
		h(w, r)
	}
	fmt.Fprintf(w, "404")
}

var service Service

// App 电影服务
func App(ctx context.Context) error {
	server := http.Server{
		Addr:    ":8080",
		Handler: &service,
	}
	go func() {
		<-ctx.Done()
		server.Shutdown(context.TODO())
	}()

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return errcode.Error{Code: errcode.ErrServerClose, Err: err}
}

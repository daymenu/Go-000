package moive

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"

	"github.com/daymenu/Go-000/Week03/work/code"
)

type moiveErr struct {
	err  error
	code code.ErrCode
}

func (err moiveErr) Error() string {
	return fmt.Sprintf("code: %d, err: %v", err.code, err.err)
}

func (err moiveErr) UnWrap() error {
	return err.err
}

// Moive 影片实体
type Moive struct {
	Name  string
	Actor string
}

// moivedService 广告
type moivedService struct {
	moives []Moive
	sync.RWMutex
}

func (m *moivedService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	action := strings.ToUpper(r.Form.Get("action"))
	switch action {
	case "GET":
		m.RLock()
		defer m.RUnlock()
		if len(m.moives) == 0 {
			fmt.Fprintf(w, "主人肾亏，还没有添加影片")
			return
		}

		l := len(m.moives)
		index := rand.Intn(l)
		moive := m.moives[index]
		fmt.Fprintf(w, "%s: %s", moive.Name, moive.Actor)
	case "POST":
		name, adWords := form.Get("name"), form.Get("actor")
		m.Lock()
		defer m.Unlock()
		mi := Moive{Name: name, Actor: adWords}
		m.moives = append(m.moives, mi)

		fmt.Fprintf(w, "添加成功")
	}
}

var moiveSer moivedService

// AppServe 启动服务
func AppServe(ctx context.Context) error {
	s := http.Server{
		Addr:    ":9090",
		Handler: &moiveSer,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
	}()

	err := s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return moiveErr{err: err, code: code.ServerErrCode}
}

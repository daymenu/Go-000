package ad

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

// Ad 广告实体
type Ad struct {
	Name    string
	ADWords string
}

// adService 广告
type adService struct {
	ads []Ad
	sync.RWMutex
}

func (ad *adService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.Form
	action := strings.ToUpper(r.Form.Get("action"))
	switch action {
	case "GET":
		ad.RLock()
		defer ad.RUnlock()
		if len(ad.ads) == 0 {
			fmt.Fprintf(w, "史莱克学院 广告招商")
			return
		}

		l := len(ad.ads)
		index := rand.Intn(l)
		moive := ad.ads[index]
		fmt.Fprintf(w, "%s: %s", moive.Name, moive.ADWords)
	case "POST":
		name, adWords := form.Get("name"), form.Get("adWords")
		ad.Lock()
		defer ad.Unlock()
		a := Ad{Name: name, ADWords: adWords}
		ad.ads = append(ad.ads, a)
		fmt.Fprintf(w, "添加成功")
	}
}

var ads adService

// AppServe 启动服务
func AppServe(ctx context.Context) error {
	s := http.Server{
		Addr:    ":9090",
		Handler: &ads,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
	}()

	err := s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return adErr{err: err, code: code.ServerErrCode}
}

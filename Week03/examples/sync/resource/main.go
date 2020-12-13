package main

import "sync"

// Resource 资源
type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}

// Resources 资源们
type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

// Poller poller
func Poller(res *Resources) {
	for {
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil {
			continue
		}
	}
}

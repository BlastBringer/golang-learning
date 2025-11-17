//load balancer

package main 

import (
	"url"
	"net/http"
	"httputil"
)

type Backend struct{
	URL *url.URL 
	Alive bool 
	mux sync.RWMutex 
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct{
	backends []*Backend 
	current uint64
}

u , _ := url.Parse("http://localhost:8000")
rp := httputil.NewSingleHostReverseProxy(u)
http.HandlerFunc(rp.ServeHTTP)


func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex() 
	l := len(s.backends) + next 

	for i := next; i <l ; i++{
		idx := i % len(s.backends)
		if s.backends[idx].isAlive(){
			if i != next {
				atomic.StoreUint64(&current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil 
}

func (b *Backend) SetAlive(alive bool){
	b.mux.Lock() 
	b.Alive = alive
	b.mux.Unlock()
}


func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock() 
	alive = b.Alive 
	b.mux.RUnlock()
	return 
}


func lb(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer() 
	if peer != nil{
		peer.ReverseProxy.ServeHTTP(w, r)
		return 
	}

	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
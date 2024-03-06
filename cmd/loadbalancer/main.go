package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type LoadBalancer struct {
  Urls []*url.URL
  Current *uint32
}

func NewLoadBalancer(urls []string) *LoadBalancer {
  servers := make([]*url.URL, len(urls))

  for i, u := range urls {
    urlParsed, err := url.Parse(u)
    if err != nil {
      panic(err)
    }
    servers[i] = urlParsed
  }

  return &LoadBalancer{
    Urls: servers,
    Current: new(uint32),
  }
}

func (lb *LoadBalancer) ReverseProxy() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    currentServer := lb.Next()
    rp := httputil.NewSingleHostReverseProxy(currentServer)
    rp.ServeHTTP(w, r)
  }
}

func (lb *LoadBalancer) Next() *url.URL {
  atomic.StoreUint32(lb.Current, uint32((*lb.Current + 1) % uint32(len(lb.Urls))))
  return lb.Urls[atomic.LoadUint32(lb.Current)]
}

func main() {
  urls := []string{"http://127.0.0.1:8081", "http://127.0.0.1:8082"}

  lb := NewLoadBalancer(urls)

  http.HandleFunc("/", lb.ReverseProxy())

  log.Fatal(http.ListenAndServe(":9999", nil))
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	URL *url.URL
	Alive bool
}

var (
	backend1, _ = url.Parse("http://localhost:8081")
	backend2, _ = url.Parse("http://localhost:8082")
	backend3, _ = url.Parse("http://localhost:8083")

	b1 = Backend{URL: backend1, Alive: true}
	b2 = Backend{URL: backend2, Alive: true}
	b3 = Backend{URL: backend3, Alive: true}
	
	backends = []*Backend{&b1, &b2, &b3}
	count int
	mu    sync.Mutex
)

func checkBackends(backends []*Backend)  {
	for _, b := range backends {
		go func(b *Backend){		  
			url := b.URL.String()
			client := http.Client{Timeout: 3 * time.Second}
			resp, err := client.Get(url)
			if err == nil {
				if resp.StatusCode == 200 {
					b.Alive = true
				}
			} else {
				b.Alive = false
			}
		}(b)
	}
}

func main()  {
	go func(){
	  for {
		checkBackends(backends)
		time.Sleep(5 * time.Second)
	  }
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		filteredList := []*Backend{}
		for _, backend := range backends {
			if backend.Alive {
				filteredList = append(filteredList, backend)
			}
		}
		
		if len(filteredList) == 0 {
			mu.Unlock()
			http.Error(w, "Service Unavailable", 503)
			return
		}

		count++
		idx := count % len(filteredList)
		server := filteredList[idx].URL
		mu.Unlock()
		
		proxy := httputil.NewSingleHostReverseProxy(server)

		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
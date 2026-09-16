package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)


var (
	backend1, _ = url.Parse("http://localhost:8081")
	backend2, _ = url.Parse("http://localhost:8082")
	backend3, _ = url.Parse("http://localhost:8083")
	
	backends = []*url.URL{backend1, backend2, backend3}
	count int
	mu    sync.Mutex
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count %= 3
		server := backends[count]
		count++
		mu.Unlock()
		
		proxy := httputil.NewSingleHostReverseProxy(server)

		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type URLStore struct {
	URLs map[string]string
	mu sync.Mutex
	counter int
}

type URL struct {
	URL string `json:"url"`
}

type Shortresponse struct {
	Short string `json:"short"`
}

func (u *URLStore) Shorten(w http.ResponseWriter, r *http.Request)  {
	u.mu.Lock()
	defer u.mu.Unlock()
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}

	defer r.Body.Close()

	data := URL{}
	err = json.Unmarshal(url, &data)
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	u.counter++
	short := strconv.Itoa(u.counter)
	u.URLs[short] = data.URL

	resp := Shortresponse{Short: short}

	res, err := json.MarshalIndent(resp, "", "\n")
	if err != nil {
		http.Error(w, "Internal error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main()  {
	store := &URLStore{
		URLs: make(map[string]string),
	}

	http.HandleFunc("/shorten", store.Shorten)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
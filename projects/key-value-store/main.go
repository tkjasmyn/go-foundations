package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

type store struct {
	data map[string]string
	mu sync.Mutex
}

type request struct {
	Value string `json:"value"`
}

type response struct {
	Status string `json:"status"`
}

func (s *store) PUT(w http.ResponseWriter, r *http.Request)  {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := strings.TrimPrefix(r.URL.Path, "/key/")

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	request := request{}
	err = json.Unmarshal(data, &request)
	if err != nil {
		fmt.Printf("Error:%v", err)
		return
	}
	s.data[key] = request.Value
	resp := response{Status: "stored"}
	
	res, err := json.MarshalIndent(resp, " ", "\n")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (s *store) GET(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/key/")

	if _, ok := s.data[key]; ok {
		value := s.data[key]
		resp := request{Value: value}

		res, err := json.MarshalIndent(resp, " ", "\n")
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		http.Error(w, "Not found", 404)
		return
	}
}

func (s *store) HandleKey(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		s.GET(w, r)
	case "PUT":
		s.PUT(w, r)
	}
}

func main()  {
	store := &store{
		data: make(map[string]string),
	}
	http.HandleFunc("/key/", store.HandleKey)
	http.ListenAndServe(":8080", nil)
}
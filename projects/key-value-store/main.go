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
		http.Error(w, err.Error(), 500)
		return
	}

	request := request{}
	err = json.Unmarshal(data, &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	s.data[key] = request.Value
	resp := response{Status: "stored"}
	s.respondJSON(w, resp)
}

func (s *store) GET(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := strings.TrimPrefix(r.URL.Path, "/key/")

	if _, ok := s.data[key]; ok {
		value := s.data[key]
		resp := request{Value: value}
		s.respondJSON(w, resp)
	} else {
		http.Error(w, "Not found", 404)
		return
	}
}

func (s *store) DELETE(w http.ResponseWriter, r *http.Request)  {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := strings.TrimPrefix(r.URL.Path, "/key/")

	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		resp := response{Status: "deleted"}
		s.respondJSON(w, resp)
	} else {
		http.Error(w, "Not found", 404)
		return
	}
}

func (s *store) respondJSON(w http.ResponseWriter, data any)  {
	res, err := json.MarshalIndent(data, " ", "\n")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (s *store) HandleKey(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		s.GET(w, r)
	case "PUT":
		s.PUT(w, r)
	case "DELETE":
		s.DELETE(w, r)
	}
}

func main()  {
	store := &store{
		data: make(map[string]string),
	}
	http.HandleFunc("/key/", store.HandleKey)
	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
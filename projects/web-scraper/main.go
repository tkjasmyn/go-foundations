package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(w http.ResponseWriter, _ *http.Request)  {
	resp, err := http.Get("https://example.com")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Body Length:", len(data))
	fmt.Println("First few bytes:", data[:10])
}

func main()  {
	http.HandleFunc("/", fetch)
	http.ListenAndServe(":8080", nil)
}
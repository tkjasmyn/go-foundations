package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetch(url string)  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("URL:", resp.Request.URL)
	fmt.Println("Body Length:", len(data))
}

func main()  {
	go fetch("http://example.com")
	go fetch("http://google.com")
	go fetch("http://facebook.com")

	time.Sleep(5 * time.Second)
}
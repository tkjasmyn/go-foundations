package main

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Println("Body Length:", len(data))
}

func main()  {
	fetch("http://example.com")
}
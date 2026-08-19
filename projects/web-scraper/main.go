package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(url string, ch chan string)  {
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

	ch <- fmt.Sprintf("URL: %s, Status: %d, Length: %d", url, resp.StatusCode, len(data))
}

func main()  {
	ch := make(chan string)
	go fetch("http://example.com", ch)
	go fetch("http://google.com", ch)
	go fetch("http://facebook.com", ch)

	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
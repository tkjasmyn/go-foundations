package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func fetch(url string, ch chan string)  {
	defer wg.Done()	

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("URL: %s, Error: %v", url, err)
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
	ch := make(chan string, 3)
	wg.Add(3)
	go fetch("http://example.com", ch)
	go fetch("http://google.com", ch)
	go fetch("http://this-does-not-exist.com", ch)
	wg.Wait()

	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
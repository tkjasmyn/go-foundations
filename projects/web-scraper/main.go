package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		ch <- fmt.Sprintf("URL: %s, Error reading body: %v", url, err)
		return
	}

	bodyString := string(data)
	start := strings.Index(bodyString, "<title>")
	end := strings.Index(bodyString, "</title>")

	if start == -1 || end == -1 {
		ch <- fmt.Sprintf("URL: %s, Title: not found", url)
		return
	}
	title := bodyString[start+len("<title>"):end]

	ch <- fmt.Sprintf("Title: %s", title)
}

func main()  {
	ch := make(chan string, 3)
	wg.Add(3)
	go fetch("http://example.com", ch)
	go fetch("http://google.com", ch)
	go fetch("http://facebook.com", ch)
	wg.Wait()

	var result []string
	for i := 0; i < 3; i++ {
		result = append(result, <-ch)
	}

	res := strings.Join(result, "\n")
	err := os.WriteFile("result.txt", []byte(res), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
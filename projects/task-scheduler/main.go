package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func printLock(mu *sync.Mutex, f func())  {
	mu.Lock()
	defer mu.Unlock()
	f()
}

func main()  {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		mu.Lock()
		fmt.Println("Task name:")
		scanner.Scan()
		mu.Unlock()
		taskName := scanner.Text()

		if taskName == "exit" {
			break
		}

		mu.Lock()
		fmt.Println("Delay:")
		scanner.Scan()
		mu.Unlock()
		delay := scanner.Text()
	
		d, err := time.ParseDuration(delay)
		if err != nil {
			fmt.Println("Enter a valid delay string")
			continue
		}	

		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(d)
			printLock(&mu, func() {
				fmt.Printf("Task %s completed\n", taskName)
			})
		}()
	}
	wg.Wait()
	printLock(&mu, func() {
		fmt.Println("All tasks completed. Exiting")
	})
}
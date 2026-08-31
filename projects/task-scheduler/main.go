package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func printLocked(mu *sync.Mutex, f func())  {
	mu.Lock()
	defer mu.Unlock()
	f()
}

func main()  {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		mu.Lock()
		fmt.Println("Welcome!")
		fmt.Println("Task Name:")
		scanner.Scan()
		taskName := scanner.Text()
		mu.Unlock()

		if taskName == "exit" {
			break
		}
		
		mu.Lock()
		fmt.Println("Delay (e.g., 5s, 2m, 10s):")
		scanner.Scan()
		delay := scanner.Text()
		mu.Unlock()
	
		d, err := time.ParseDuration(delay)
		if err != nil {
			printLocked(&mu, func() {
				fmt.Println("Enter a valid delay string")
			})
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(d)
			printLocked(&mu, func() {
				fmt.Printf("Task %s executed after %d seconds\n", taskName, int(d.Seconds()))
			})
			}()
		}
		wg.Wait()

		printLocked(&mu, func() {
			fmt.Println("All tasks completed. Exiting.")
		})
	}
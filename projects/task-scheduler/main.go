package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type Task struct {
	ID int
	Name string
	Delay string
}

func printLocked(mu *sync.Mutex, f func())  {
	mu.Lock()
	defer mu.Unlock()
	f()
}

func main()  {
	var (
		nextID int
		pending []Task
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

		if taskName == "list" {
				fmt.Println("Pending tasks:")
			for i, task := range pending {
				fmt.Printf("  %d. %s (%s)\n", i+1, task.Name, task.Delay)
			}
				continue
			}
		
		mu.Lock()
		fmt.Println("Delay (e.g., 5s, 2m, 10s):")
		scanner.Scan()
		delay := scanner.Text()
		nextID++
		task := Task{
			ID: nextID,
			Name: taskName,
			Delay: delay,
		}
		pending = append(pending, task)
		fmt.Printf("Task %d scheduled: %s (%s)\n", nextID, taskName, delay)
		mu.Unlock()
	
		d, err := time.ParseDuration(delay)
		if err != nil {
			printLocked(&mu, func() {
				fmt.Println("Enter a valid delay string")
			})
			continue
		}

		wg.Add(1)
		id := nextID
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(d)
			printLocked(&mu, func() {
				fmt.Printf("Task %s executed after %d seconds\n", taskName, int(d.Seconds()))
				for i := range pending {
					if pending[i].ID == taskID {
						pending = append(pending[:i], pending[i+1:]...)
						break
					}
				}
			})
			}(id)
		}
		wg.Wait()

		printLocked(&mu, func() {
			fmt.Println("All tasks completed. Exiting.")
		})
	}
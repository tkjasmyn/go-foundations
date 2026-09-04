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

func printLock(mu *sync.Mutex, f func())  {
	mu.Lock()
	defer mu.Unlock()
	f()
}

func printList(mu *sync.Mutex, pending []Task)  {
	mu.Lock()
	defer mu.Unlock()

	if len(pending) == 0 {
		fmt.Println("No pending tasks")
		return
	}

	fmt.Println("Pending tasks:")
	for i, task := range pending {
		fmt.Printf("  %d. %s (%s)\n", i+1, task.Name, task.Delay)
	}
}

func main()  {
	var (
		nextID int
		mu sync.Mutex
		wg sync.WaitGroup
	)
	pending := []Task{}

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

		if taskName == "list" {
			printList(&mu, pending)
			continue
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
		nextID++

		pending = append(pending, Task{ID: nextID, Name: taskName, Delay: delay})

		id := nextID
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(d)
			printLock(&mu, func() {
				fmt.Printf("Task %s completed\n", taskName)
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
	printLock(&mu, func() {
		fmt.Println("All tasks completed. Exiting")
	})
}
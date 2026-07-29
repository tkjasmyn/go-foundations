package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
}

type TaskStore struct {
	tasks []Task
	nextID int
	filename string
}

func (ts *TaskStore) loadTasks()  {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	var task []Task

	err = json.Unmarshal(data, &task)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	ts.tasks = append(ts.tasks, task...)
	ts.nextID = 1
	for _, task := range ts.tasks {
	    if task.ID >= ts.nextID {
	        ts.nextID = task.ID + 1
	    }
	}
}

func (ts *TaskStore) saveTasks()  {
	data, err := json.MarshalIndent(ts.tasks, " ", "\t")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	err = os.WriteFile(ts.filename, data, 0644)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}

func (ts *TaskStore) add(desc string)  {
	task := Task {
		ID: ts.nextID,
		Description: desc,
		Status: "todo",
	}

	ts.tasks = append(ts.tasks, task)
	ts.nextID++
}

func (ts *TaskStore) list()  {
	if len(ts.tasks) == 0 {
		fmt.Println("No tasks yet.")
	}
	for _, task := range ts.tasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func main()  {
	store := &TaskStore {
		tasks: []Task{},
		nextID: 1,
		filename: "tasks.json",
	}
	store.loadTasks()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What do you want to do? (add/list/exit)")
	scanner.Scan()
	if scanner.Text() == "add" {
		fmt.Println("Enter task description:")
		scanner.Scan()
		desc := scanner.Text()
		store.add(desc)
		fmt.Println("Task added successfully")
		store.saveTasks()
	} else if scanner.Text() == "list" {
		store.list()
	}
}
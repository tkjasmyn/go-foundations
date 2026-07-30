package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if os.IsNotExist(err) {
			return
		}
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

func (ts *TaskStore) done(id int)  {
	found := false
	for i := range ts.tasks {
		if ts.tasks[i].ID == id {
			ts.tasks[i].Status = "done"
			fmt.Printf("Task %d marked done\n",ts.tasks[i].ID)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found.")
	}
}

func (ts *TaskStore) delete(id int)  {
	found := false
	for i := range ts.tasks {
		if ts.tasks[i].ID == id {
			fmt.Printf("Task %d deleted successfully.\n", id)
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found.")
	}
}

func (ts *TaskStore) deleteAll()  {
	err := os.Remove("./tasks.json")
	if err != nil {
		fmt.Println("No file to be deleted")
		return
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

	for {
		fmt.Println("What do you want to do? (add/list/exit)")
		scanner.Scan()
		command := scanner.Text()
		if command == "add" {
			fmt.Println("Enter task description:")
			scanner.Scan()
			desc := scanner.Text()
			store.add(desc)
			fmt.Println("Task added successfully")
			store.saveTasks()
		} else if command == "list" {
			store.list()
		} else if command == "done" {
			fmt.Println("Enter task id:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("Error: %v\nPlease enter a valid number\n", err)
				return
			}
			store.done(id)
			store.saveTasks()
		} else if command == "delete" {
			fmt.Println("Enter task id:")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("Error: %v\nPlease enter a valid number\n", err)
				return
			}
			store.delete(id)
			store.saveTasks()
		} else if command == "delete all" {
			store.deleteAll()
			fmt.Println("File deleted")
		} else if command == "exit" {
			fmt.Println("Goodbye")
			break
		}
	}
}
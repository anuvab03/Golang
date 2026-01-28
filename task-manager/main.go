package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task[add|list|done|delete]")
		return
	}

	command := os.Args[1]
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Task title required")
			return
		}

		title := os.Args[2]
		id := len(tasks) + 1
		tasks = append(tasks, Task{ID: id, Title: title})
		saveTasks(tasks)

	case "list":
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s]%d:%s\n", status, t.ID, t.Title)
		}
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				break
			}
		}
		saveTasks(tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Task ID required")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		var updated []Task
		for _, t := range tasks {
			if t.ID != id {
				updated = append(updated, t)
			}
		}
		saveTasks(updated)

	default:
		fmt.Println("Unknown Command")
	}
}

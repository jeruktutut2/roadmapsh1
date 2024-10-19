package main

import (
	"fmt"
	"os"
)

type TaskTracker struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	filename := "task_tracker.json"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("TASK TRACKER")
	fmt.Println(`# Adding a new task: add "Buy groceries"`)
	fmt.Println(`# Updating tasks: update 1 "Buy groceries and cook dinner"`)
	fmt.Println(`# Deleting tasks: delete 1`)
	fmt.Println(`# Marking a task as in progress: mark-in-progress 1`)
	fmt.Println(`# Marking a task as in done: mark-done 1`)
	fmt.Println(`# Listing all tasks: list`)
	fmt.Println(`# Listing tasks by status done: list done`)
	fmt.Println(`# Listing tasks by status to do: list todo`)
	fmt.Println(`# Listing tasks by status in progress: list in-progress`)
	fmt.Println(`# To exit: exit`)

	for {
		var command string
		fmt.Print("command:")
		fmt.Scanln(&command)
		// fmt.Println("command:", command)
		switch command {
		case "exit":
			fmt.Println("Thank You")
			os.Exit(0)
		}
	}
}

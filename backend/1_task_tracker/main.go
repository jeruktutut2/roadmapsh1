package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskTracker struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
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

taskLoop:
	for {
		var task string
		reader := bufio.NewReader(os.Stdin)
		task, err := reader.ReadString('\n')
		if err != nil {
			panic("error when read string from reader: " + err.Error())
		}
		task = strings.TrimSpace(task)
		commands := strings.Split(task, " ")
		command := commands[0]
		switch command {
		case "add":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			startIndex := strings.Index(task, `"`)
			lastIndex := strings.LastIndex(task, `"`)
			description := task[startIndex+1 : lastIndex]
			var id int
			if len(taskTrackers) > 0 {
				taskTracker := taskTrackers[len(taskTrackers)-1]
				id = taskTracker.Id + 1
			} else {
				id = 1
			}

			var taskTracker TaskTracker
			taskTracker.Id = id
			taskTracker.Description = description
			taskTracker.Status = "todo"
			taskTracker.CreatedAt = time.Now().String()
			taskTrackers = append(taskTrackers, taskTracker)

			err = writeToFile(file, taskTrackers)
			if err != nil {
				fmt.Println("error when writing to file:", err)
				continue taskLoop
			}
			continue taskLoop
		case "update":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("error when converting id from int to string:", err)
				continue taskLoop
			}
			startIndex := strings.Index(task, `"`)
			lastIndex := strings.LastIndex(task, `"`)
			description := task[startIndex+1 : lastIndex]
			isTaskExists := false
		taskTrackerUpdateLoop:
			for i := range taskTrackers {
				if taskTrackers[i].Id == id {
					taskTrackers[i].Description = description
					taskTrackers[i].UpdatedAt = time.Now().String()
					isTaskExists = true
					break taskTrackerUpdateLoop
				}
			}
			if !isTaskExists {
				fmt.Println("cannot find data with id:", id)
				continue taskLoop
			}
			err = writeToFile(file, taskTrackers)
			if err != nil {
				fmt.Println("error when writing to file:", err)
				continue taskLoop
			}
		case "delete":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("error when converting id from int to string:", err)
				continue taskLoop
			}
			isTaskExists := false
		taskTrackerDeleteLoop:
			for i := range taskTrackers {
				if taskTrackers[i].Id == id {
					taskTrackers = append(taskTrackers[:i], taskTrackers[i+1:]...)
					isTaskExists = true
					break taskTrackerDeleteLoop
				}
			}
			if !isTaskExists {
				fmt.Println("cannot find data with id:", id)
				continue taskLoop
			}
			err = writeToFile(file, taskTrackers)
			if err != nil {
				fmt.Println("error when writing to file:", err)
				continue taskLoop
			}
		case "mark-in-progress":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("error when converting id from int to string:", err)
				continue taskLoop
			}
			isTaskExists := false
		taskTrackerMarkInProgressLoop:
			for i := range taskTrackers {
				if taskTrackers[i].Id == id {
					taskTrackers[i].Status = "in-progress"
					isTaskExists = true
					break taskTrackerMarkInProgressLoop
				}
			}
			if !isTaskExists {
				fmt.Println("cannot find data with id:", id)
				continue taskLoop
			}
			err = writeToFile(file, taskTrackers)
			if err != nil {
				fmt.Println("error when writing to file:", err)
				continue taskLoop
			}
		case "mark-done":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			id, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("error when converting id from int to string:", err)
				continue taskLoop
			}
			isTaskExists := false
		taskTrackerMarkDoneLoop:
			for i := range taskTrackers {
				if taskTrackers[i].Id == id {
					taskTrackers[i].Status = "done"
					isTaskExists = true
					break taskTrackerMarkDoneLoop
				}
			}
			if !isTaskExists {
				fmt.Println("cannot find data with id:", id)
				continue taskLoop
			}
			err = writeToFile(file, taskTrackers)
			if err != nil {
				fmt.Println("error when writing to file:", err)
			}
		case "list":
			file, taskTrackers, err := openAndReadFile()
			if err != nil {
				fmt.Println("error when open and read file:", err)
				continue taskLoop
			}
			defer file.Close()
			if len(commands) < 2 {
				for _, taskTracker := range taskTrackers {
					fmt.Println(taskTracker)
				}
				continue taskLoop
			} else {
				switch commands[1] {
				case "done":
					for _, taskTracker := range taskTrackers {
						if taskTracker.Status == "done" {
							fmt.Println(taskTracker)
						}
					}
					continue taskLoop
				case "todo":
					for _, taskTracker := range taskTrackers {
						if taskTracker.Status == "todo" {
							fmt.Println(taskTracker)
						}
					}
					continue taskLoop
				case "in-progress":
					for _, taskTracker := range taskTrackers {
						if taskTracker.Status == "in-progress" {
							fmt.Println(taskTracker)
						}
					}
					continue taskLoop
				default:
					fmt.Println("please choose done, todo or in-progress status")
					continue taskLoop
				}
			}
		case "exit":
			fmt.Println("Thank You")
			os.Exit(0)
		default:
			fmt.Println("wrong command")
			continue taskLoop
		}
	}
}

func openAndReadFile() (file *os.File, taskTracekers []TaskTracker, err error) {
	filename := "task_tracker.json"
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	info, err := file.Stat()
	if err != nil {
		panic("error when getting file information: " + err.Error())
	}
	if info.Size() == 0 {
		return
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&taskTracekers); err != nil && err != io.EOF {
		return
	}
	// file.Seek(0, 0) will place the pointer to first line and column (let's say like that), so if already have data that more than new data, the last data will not be replaced, the last data will be still there
	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}
	return
}

func writeToFile(file *os.File, taskTrackers []TaskTracker) (err error) {
	taskTrackersByte, err := json.Marshal(taskTrackers)
	if err != nil {
		return
	}
	err = os.Truncate(file.Name(), 0)
	if err != nil {
		return
	}
	_, err = file.WriteString(string(taskTrackersByte))
	if err != nil {
		return
	}
	return
}

package commands

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func handleAdd(args []string, file *os.File) error {
	if len(args) < 1 {
		return fmt.Errorf("please add a description of the task")
	}

	tasks, err := getTasks(file)

	if err != nil {
		return err
	}

	var lastID int = 0

	if len(*tasks) != 0 {
		lastTask := (*tasks)[len(*tasks)-1]
		lastID = lastTask.ID
	}

	task := createTask(args, lastID)
	*tasks = append(*tasks, task)

	if err = saveTasks(file, tasks); err != nil {
		return err
	}

	fmt.Printf("task added successfully (ID: %d)\n", task.ID)

	return nil
}

func createTask(args []string, lastID int) task {
	now := time.Now().Format("2006-01-02 15:04:05")

	return task{
		ID:          lastID + 1,
		Description: strings.Join(args, " "),
		Status:      StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

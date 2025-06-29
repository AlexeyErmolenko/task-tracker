package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	Name        string
	Description string
	Callback    func([]string, *os.File) error
}

type task struct {
	ID          int
	Description string
	Status      int
	CreatedAt   string
	UpdatedAt   string
}

const (
	StatusToDo = iota
	StatusInProgress
	StatusDone
)

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    handleHelp,
		},
		"add": {
			Name:        "add",
			Description: "Add new task",
			Callback:    handleAdd,
		},
		"update": {
			Name:        "update",
			Description: "Update a task",
			Callback:    handleUpdate,
		},
		"delete": {
			Name:        "delete",
			Description: "Delete a task",
			Callback:    handleDelete,
		},
	}
}

func getTasks(file *os.File) (*[]task, error) {
	var tasks []task
	decoded := json.NewDecoder(file)
	err := decoded.Decode(&tasks)

	if err != nil {
		return &tasks, err
	}

	return &tasks, nil
}

func saveTasks(file *os.File, tasks *[]task) error {
	jsonData, err := json.Marshal(*tasks)

	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	_, err = file.Write(jsonData)

	return err
}

func getID(args []string) (int, error) {
	if len(args) < 1 {
		return 0, fmt.Errorf("please input ID of the task")
	}

	id, err := strconv.Atoi(args[0])

	if err != nil {
		return 0, fmt.Errorf("task ID should be a number")
	}

	return id, nil
}

package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleUpdate(args []string, file *os.File) error {
	if len(args) < 1 {
		return fmt.Errorf("please input ID of the task")
	}

	id, err := strconv.Atoi(args[0])

	if err != nil {
		return fmt.Errorf("task ID should be a number")
	}

	if len(args) < 2 {
		return fmt.Errorf("please add a description of the task")
	}

	description := strings.Join(args[1:], " ")
	tasks, err := getTasks(file)

	if err != nil {
		return err
	}

	isUpdated := false

	for i, v := range *tasks {
		if v.ID != id {
			continue
		}

		(*tasks)[i].Description = description
		(*tasks)[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		isUpdated = true
		break
	}

	if !isUpdated {
		return fmt.Errorf("task with (ID %d) not found", id)
	}

	if err = saveTasks(file, tasks); err != nil {
		return err
	}

	fmt.Printf("task with (ID %d) was updated successfully\n", id)
	return nil
}

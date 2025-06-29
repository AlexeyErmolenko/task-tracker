package commands

import (
	"fmt"
	"os"
)

func handleDelete(args []string, file *os.File) error {
	id, err := getID(args)

	if err != nil {
		return err
	}

	tasks, err := getTasks(file)

	if err != nil {
		return err
	}

	isDeleted := false

	for i, v := range *tasks {
		if v.ID != id {
			continue
		}

		*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
		isDeleted = true
		break
	}

	if !isDeleted {
		return fmt.Errorf("task with (ID %d) not found", id)
	}

	if err = saveTasks(file, tasks); err != nil {
		return err
	}

	fmt.Printf("task with (ID %d) was deleted successfully\n", id)

	return nil
}

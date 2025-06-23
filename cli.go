package main

import (
	"fmt"
	"os"

	"github.com/AlexeyErmolenko/task-tracker/commands"
)

type config struct {
	fileName string
}

func handle(conf *config) {
	args := os.Args
	command, ok := getCommand(args)

	if !ok {
		fmt.Println("Command not found")
		return
	}

	file, err := getFile(conf.fileName)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	command.Callback(args, file)
}

func getCommand(args []string) (commands.Command, bool) {
	if len(args) < 2 {
		return commands.Command{}, false
	}

	commannds := commands.GetCommands()
	command, ok := commannds[args[1]]

	return command, ok
}

func getFile(filePath string) (*os.File, error) {
	if !fileExists(filePath) {
		file, err := os.Create(filePath)

		if err != nil {
			return &os.File{}, err
		}

		_, err = file.Write([]byte("{}"))

		if err != nil {
			return file, err
		}

		file.Close()
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)

	return file, err
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

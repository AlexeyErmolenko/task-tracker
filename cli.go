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

	command.Callback(args)
}

func getCommand(args []string) (commands.Command, bool) {
	if len(args) < 2 {
		return commands.Command{}, false
	}

	commannds := commands.GetCommands()
	command, ok := commannds[args[1]]

	return command, ok
}

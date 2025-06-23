package commands

import "os"

type Command struct {
	Name        string
	Description string
	Callback    func([]string, *os.File) error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    handleHelp,
		},
	}
}

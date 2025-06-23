package commands

import "fmt"

func handleHelp(args []string) error {
	showHelp()

	return nil
}

func showHelp() {
	fmt.Println("This tool allows storing and managing your tasks.")
	fmt.Println("It has next commands:")

	commands := GetCommands()

	for _, v := range commands {
		fmt.Println("  ", v.Name, "-", v.Description)
	}
}

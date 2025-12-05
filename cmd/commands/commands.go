package commands

import (
	"github.com/Kuroashi1995/go-pokedex/config"
)

// Define cliCommand Struct
type CliCommand struct {
	Name		string
	Description	string
	Callback	func(*config.Config) error
}


func GetCommands() map[string]CliCommand {
	command_list := map[string]CliCommand{
		"exit": {
			Name: "exit",
			Description: "Exit the Pokedex",
			Callback: commandExit,
		},
		"help": {
			Name: "help",
			Description: "Displays a help message",
			Callback: commandHelp,
		},
		"map": {
			Name: "map",
			Description: "Displays the next 20 locations",
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Displays the previous 20 locations",
			Callback: commandMapBack,
		},
		"explore": {
			Name: "explore",
			Description: "Displays location's pokemon",
			Callback: commandExplore,
		},
}
	return command_list
}


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
		"catch": {
			Name: "catch",
			Description: "Try to catch the pokemon",
			Callback: commandCatch,
		},
		"inspect": {
			Name: "inspect",
			Description: "Inspect catched pokemon",
			Callback: commandInspect,
		},
		"pokedex": {
			Name: "pokedex",
			Description: "List all catched pokemon",
			Callback: commandPokedex,
		},
}
	return command_list
}


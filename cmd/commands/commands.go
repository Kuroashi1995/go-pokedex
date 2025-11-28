package commands

// Define cliCommand Struct
type CliCommand struct {
	Name		string
	Description	string
	Callback	func() error
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
}
	return command_list
}


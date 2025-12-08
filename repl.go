package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kuroashi1995/go-pokedex/cmd/commands"
	"github.com/Kuroashi1995/go-pokedex/cmd/input"
	"github.com/Kuroashi1995/go-pokedex/config"
)

func repl(cfg *config.Config) {

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user_input := scanner.Text()
		cleaned := input.CleanInput(user_input)
		for key, command := range commands.GetCommands(){
			if cleaned[0] == key{
				switch key {
				case "explore":
						cfg.CurrentLocation = &cleaned[1]
				case "catch":
					cfg.Catching = &cleaned[1]
				case "inspect":
					cfg.Inspecting = &cleaned[1]
				}
				err := command.Callback(cfg)
				if err != nil {
					fmt.Println("An error ocurred processing the command:", err)
				}
			}
		}
	}
}

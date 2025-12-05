package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kuroashi1995/go-pokedex/cmd/commands"
	"github.com/Kuroashi1995/go-pokedex/cmd/input"
	"github.com/Kuroashi1995/go-pokedex/config"
	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

func repl(cfg *config.Config, cache *pokecache.Cache) {

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user_input := scanner.Text()
		cleaned := input.CleanInput(user_input)
		for key, command := range commands.GetCommands(){
			if cleaned[0] == key{
				err := command.Callback(cfg, cache)
				if err != nil {
					fmt.Println("An error ocurred processing the command:", err)
				}
			}
		}
	}
}

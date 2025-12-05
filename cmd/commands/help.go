package commands

import (
	"errors"
	"fmt"

	"github.com/Kuroashi1995/go-pokedex/config"
	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

func commandHelp(config *config.Config, cache *pokecache.Cache) error {
	commands := GetCommands()
	if len(commands) == 0  || commands == nil {
		return errors.New("No commands")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range commands {
		cmd_string := fmt.Sprintf("%v: %v", command.Name , command.Description)
		fmt.Println(cmd_string)
	}
	return nil
}

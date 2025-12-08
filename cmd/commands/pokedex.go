package commands

import (
	"fmt"

	"github.com/Kuroashi1995/go-pokedex/config"
)

func commandPokedex(config *config.Config) error {
	//check if there are pokemon
	if len(*config.Pokedex) == 0 {
		fmt.Println("Go catch some pokemon!")
		return nil
	}


	fmt.Println("Your Pokedex:")
	for pokemon := range *config.Pokedex {
		fmt.Println(" - " + pokemon)
	}
	return nil
}

package commands

import (
	"errors"
	"fmt"

	"github.com/Kuroashi1995/go-pokedex/config"
)

func commandInspect(config *config.Config) error {
	currentlyInspecting := config.Inspecting
	pokemon, ok := (*config.Pokedex)[*currentlyInspecting]
	if !ok {
		return errors.New("Pokemon hasn't been captured yet")
	}
	//print stats
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf(" -%s\n", pokemonType.Type.Name)
	}
	return nil

}

package commands

import (
	"fmt"

	"github.com/Kuroashi1995/go-pokedex/config"
)

func commandExplore(config *config.Config) error {
	pokemonData, err := config.Client.LocationPokemons(*config.CurrentLocation, config.Cache)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonData.PokemonEncounters{
		fmt.Println("- " + pokemon.Pokemon.Name)
	}
	return nil

}

package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Kuroashi1995/go-pokedex/config"
)


func commandCatch(cfg *config.Config) error {
	//print initial message
	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.Catching)
	//get pokemon
	catching, err := cfg.Client.GetPokemon(*cfg.Catching, cfg.Cache)
	if err != nil {
		return err
	}
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	threshold := 1 - float32(650 - catching.BaseExperience) / float32(650)
	value := rand.Float32()
	if value < threshold {
		fmt.Printf("%s escaped!\n", catching.Name)
		return nil
	} else {
		(*cfg.Pokedex)[catching.Name] = catching
		fmt.Printf("%s was caught!\n", catching.Name)
		fmt.Printf("You may now inspect it with the inspect command\n")
		return nil
	}
}

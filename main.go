package main

import (
	"time"

	"github.com/Kuroashi1995/go-pokedex/config"
	"github.com/Kuroashi1995/go-pokedex/internal/pokeapi"
	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.GetClient(time.Second * 5)
	pokeCache := pokecache.NewCache(time.Second * 30)
	pokedex := make(map[string]pokeapi.RespPokemon)
	cfg := &config.Config{
		Client: pokeClient,
		Cache: pokeCache,
		Pokedex: &pokedex,
	}
	repl(cfg)
}

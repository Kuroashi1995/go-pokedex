package config

import "github.com/Kuroashi1995/go-pokedex/internal/pokeapi"

type Config struct {
	Client	pokeapi.Client
	Prev	*string
	Next	*string
}

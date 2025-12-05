package config

import (
	"github.com/Kuroashi1995/go-pokedex/internal/pokeapi"
	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

type Config struct {
	Client			pokeapi.Client
	Cache			*pokecache.Cache
	CurrentLocation	*string
	Prev			*string
	Next			*string
}

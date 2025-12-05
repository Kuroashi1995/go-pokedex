package commands

import (
	"fmt"
	"os"

	"github.com/Kuroashi1995/go-pokedex/config"
	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

// Define exit command
func commandExit(config *config.Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

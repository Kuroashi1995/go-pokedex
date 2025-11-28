package commands

import (
	"os"
	"fmt"
)
// Define exit command
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

package commands

import (
	"fmt"

	"github.com/Kuroashi1995/go-pokedex/config"
)


func commandMap(config *config.Config) error {

	locationsData, err := config.Client.ListLocations(config.Next, config.Cache)
	if err != nil {
		return err
	}
	//set config values
	config.Next, config.Prev = locationsData.Next, locationsData.Previous

	for _, location := range locationsData.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandMapBack(config *config.Config) error {
	if config.Prev== nil{
		fmt.Println("you're on the first page")
		return nil
	}
	// Initialize request
	locationsData, err := config.Client.ListLocations(config.Prev, config.Cache)
	if err != nil {
		return err
	}

	//set config values
	config.Next, config.Prev = locationsData.Next, locationsData.Previous

	for _, location := range locationsData.Results {
		fmt.Println(location.Name)
	}
	return nil
}

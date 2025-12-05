package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)



func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespLocations, error) {
	// Initialize request
	url := baseURL + "/location-area"
	// Why does he do this?
	if pageURL != nil {
		url = *pageURL
	}
	cache_data, ok := cache.Get(url)
	if ok {
		var cache_hit RespLocations
		if err := json.Unmarshal(cache_data, &cache_hit); err != nil {
			return RespLocations{}, err
		}
		return cache_hit, nil
	}
	request, err := http.NewRequest("GET", url,  nil)
	if err != nil {
		fmt.Println("an error ocurred initializing the request", err)
		return RespLocations{}, err
	}

	// Make request
	res, err := c.httpClient.Do(request)
	if err != nil  {
		fmt.Println("an error ocurred making the request", err)
		return RespLocations{}, err
	}
	defer res.Body.Close()

	// Unmarshal data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}
	cache.Add(url, data)
	var locations_data RespLocations
	if err := json.Unmarshal(data, &locations_data); err != nil {
		fmt.Println("an error ocurred unmarhaslling the request", err)
		return RespLocations{}, err
	}
	return locations_data, nil
}

func (c *Client) LocationPokemons(locationName string, cache *pokecache.Cache) (RespLocationPokemon, error){
	// configure url
	url := baseURL + "/location-area/" + locationName + "/"
	if locationName == "" {
		return RespLocationPokemon{}, errors.New("Location name can't be empty")
	}
	// check cache
	cache_data, ok := cache.Get(url)
	if ok {
		var cache_hit RespLocationPokemon
		if err := json.Unmarshal(cache_data, &cache_hit); err != nil {
			return RespLocationPokemon{}, err
		}
		return cache_hit, nil
	}
	// prepare request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPokemon{}, err
	}
	// make request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPokemon{}, err
	}
	defer res.Body.Close()
	//process response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationPokemon{}, err
	}
	// add to cache
	cache.Add(url, data)
	var pokemon_data RespLocationPokemon
	if err := json.Unmarshal(data, &pokemon_data); err != nil {
		return RespLocationPokemon{}, err
	}
	return pokemon_data, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kuroashi1995/go-pokedex/internal/pokecache"
)

func (c *Client) GetPokemon(pokemonName string, cache *pokecache.Cache) (RespPokemon, error) {
	// set url
	url := baseURL + "/pokemon/" + pokemonName + "/"
	//check cache
	cacheHit, ok := cache.Get(url)
	if ok {
		var cacheData RespPokemon
		if err := json.Unmarshal(cacheHit, &cacheData); err != nil {
			return RespPokemon{}, err
		}
		return cacheData, nil
	}
	// set up request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("An error ocurred while initializing the request")
		return RespPokemon{}, err
	}
	//make request
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("An error ocurred while making the request")
		return RespPokemon{}, err
	}
	defer res.Body.Close()
	//process request body
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("An error ocurred while processing body data")
		return RespPokemon{}, err
	}
	//save cache
	cache.Add(url, resData)
	//unmarshall and return data
	var pokemonData RespPokemon 
	if err := json.Unmarshal(resData, &pokemonData); err != nil {
		fmt.Println("An error ocurred while unmarhaslling data")
		return RespPokemon{}, err
	}
	return pokemonData, nil
}


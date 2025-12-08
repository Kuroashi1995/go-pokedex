package pokeapi

type RespLocations struct {
	Count		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]struct {
		Name	string		`json:"name"`
		URL		string		`json:"url"`
	}						`json:"results"`
}

type RespLocationPokemon struct {
	Id						int		`json:"id"`
	Name					string	`json:"name"`
	PokemonEncounters		[]struct{
		Pokemon		struct{
			Name	string			`json:"name"`
			Url		string			`json:"url"`
		}							`json:"pokemon"`
	}								`json:"pokemon_encounters"`
}

type RespPokemon struct {
	Id						int		`json:"id"`
	Name					string	`json:"name"`
	BaseExperience			int		`json:"base_experience"`
	Height					int		`json:"height"`
	Weight					int		`json:"weight"`
	Order					int		`json:"order"`
	Stats					[]struct {
		BaseStat			int		`json:"base_stat"`
		Stat				struct {
			Name			string	`json:"name"`
			Url				string	`json:"url"`
		}							`json:"stat"`
	}								`json:"stats"`
	Types					[]struct {
		Slot				int		`json:"slot"`
		Type				struct {
			Name			string	`json:"name"`
			Url				string	`json:"url"`
		}							`json:"type"`
	}								`json:"Types"`
}

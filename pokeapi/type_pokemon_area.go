package pokeapi

type PokemonDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEntry struct {
	Pokemon PokemonDetail `json:"pokemon"`
}

type PokemonEncounter struct {
	PokemonEncounter []PokemonEntry `json:"pokemon_encounters"`
}

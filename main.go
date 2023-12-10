package main

import (
	"time"
	"github.com/AnhTuan4198/Pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(20 * time.Second);
	config := &config{
		pokeApiClient: pokeClient,
	}
	startRELP(config);
}

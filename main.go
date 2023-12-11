package main

import (
	"time"

	"github.com/AnhTuan4198/Pokedex/pokeapi"
	"github.com/AnhTuan4198/Pokedex/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(20 * time.Second);
	cache := pokecache.NewCache(10 * time.Second);
	config := &config{
		pokeApiClient: pokeClient,
		cache: cache,
	}
	startRELP(config);
}

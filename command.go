package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AnhTuan4198/Pokedex/pokeapi"
	"github.com/AnhTuan4198/Pokedex/pokecache"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")
	for _, val := range getCommands() {
		val.getInformation()
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
	return errors.New("error")
}

func commandMap(cfg *config, args ...string) error {
	var res = pokeapi.LocationAreaResponse{}
	hasCacheData := false;
	cacheData := pokecache.CacheEntry{};
	if cfg.nextLocationUrl != nil{
		data, ok := cfg.cache.Get(*cfg.nextLocationUrl);
		hasCacheData = ok;
		cacheData = data;
	}

	fmt.Println(cacheData);
	
	if !hasCacheData {
		resp, err := cfg.pokeApiClient.GetLocationList(cfg.nextLocationUrl, cfg.cache)
		if err != nil {
			return nil
		}
		res = resp;
	} else {
		ok := json.Unmarshal(cacheData.Value, &res);
		if ok != nil {
			return nil
		}
	}

	cfg.nextLocationUrl = res.Next
	cfg.prevLocationUrl = res.Previous

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}

func commandMapb(cfg *config,  args ...string) error {
	if cfg.prevLocationUrl == nil {
		fmt.Println("Unable to go back!")
		return nil
	}
	var res = pokeapi.LocationAreaResponse{}
	var nextUrl *string
	var preUrl *string
	cacheData, ok := cfg.cache.Get(*cfg.prevLocationUrl)
	if !ok {
		fmt.Println("Get new prev localtions")
		res, err := cfg.pokeApiClient.GetLocationList(cfg.prevLocationUrl, cfg.cache)
		if err != nil {
			return nil
		}
		nextUrl = res.Next
		preUrl = res.Previous
	} else {
		ok := json.Unmarshal(cacheData.Value, &res)
		if ok != nil {
			return nil
		}
	}

	cfg.nextLocationUrl = nextUrl
	cfg.prevLocationUrl = preUrl

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}


func commandExplore(cfg *config, args ...string) error {
	// var res = pokeapi.PokemonEncounter{}
	locationName := args[0]
	if len(locationName) == 0{
		return errors.New("Invalid location")
	}
	fmt.Println("Exploring pastoria-city-area...")
	// var res = pokeapi.PokemonEncounter{};

	pokemonEncounter, error := cfg.pokeApiClient.ExplorePokemonInArea(locationName, cfg.cache);

	if error != nil{
		fmt.Println("Not found!")
		return errors.New("Something when wrong")
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonEncounter.PokemonEncounter{
		fmt.Println("-",pokemon.Pokemon.Name);
	}

	fmt.Println(args);
	return nil
}

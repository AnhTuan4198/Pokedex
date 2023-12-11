package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AnhTuan4198/Pokedex/pokeapi"
	"github.com/AnhTuan4198/Pokedex/pokecache"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")
	for _, val := range getCommands() {
		val.getInformation()
	}
	return nil
}

func commandExit(cfg *config) error {
	return errors.New("error")
}

func commandMap(cfg *config) error {
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

	fmt.Println("Next", *res.Next, res.Previous);

	for _, item := range res.Results {
		fmt.Println(item.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
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

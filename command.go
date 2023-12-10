package main

import (
	"errors"
	"fmt"
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

func commandExit(cfg * config) error {
	return errors.New("error")
}

func commandMap(cfg * config) error {
	res, err := cfg.pokeApiClient.GetLocationList(cfg.nextLocationUrl);
	if err != nil{
		return err
	}

	cfg.nextLocationUrl = res.Next
	cfg.prevLocationUrl = res.Previous

	for _, item := range res.Results{
		fmt.Println(item.Name);
	}
	return nil
}

func commandMapb(cfg * config) error {
	if cfg.prevLocationUrl == nil {
		fmt.Println("Unable to go back!");
		return nil;
	}
	res, err := cfg.pokeApiClient.GetLocationList(cfg.prevLocationUrl);
	if err != nil{
		return err
	}

	cfg.nextLocationUrl = res.Next
	cfg.prevLocationUrl = res.Previous

	for _, item := range res.Results{
		fmt.Println(item.Name);
	}
	return nil
}


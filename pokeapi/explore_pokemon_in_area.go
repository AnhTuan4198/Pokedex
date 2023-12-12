package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/AnhTuan4198/Pokedex/pokecache"
)

func (c *Client) ExplorePokemonInArea(area_name string, cache *pokecache.Cache) (PokemonEncounter, error) {
	if len(area_name) == 0 {
		return PokemonEncounter{}, nil
	}
	url := baseURL + "/location-area/" + area_name;
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounter{}, errors.New("Something wrong")
	}
	resp, err := c.httpClient.Do(req);
	if err != nil {
		return PokemonEncounter{}, errors.New("Explore pokemon fail")
	}
	defer resp.Body.Close()
	exploreData := PokemonEncounter{}
	dataByte, err := io.ReadAll(resp.Body);
	if err != nil {
		return PokemonEncounter{}, errors.New("Read data from response fail")
	}
	parseJsonErr := json.Unmarshal(dataByte, &exploreData)
	if parseJsonErr != nil {
		return PokemonEncounter{}, errors.New("Fail to parse explore JSON")
	}

	cfgCache := *cache
	cfgCache.Add(url, dataByte)
	return exploreData, nil
}

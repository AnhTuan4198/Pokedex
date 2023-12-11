package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AnhTuan4198/Pokedex/pokecache"
)

func (c *Client) GetLocationList(pageURl *string, cache *pokecache.Cache) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"

	if pageURl != nil {
		url = *pageURl
	}

	// create request
	req, error := http.NewRequest("GET",url, nil);
	if error != nil {
		return LocationAreaResponse{}, error
	}
	fmt.Println(url)
	response, err := c.httpClient.Do(req);
	if err != nil{
		return LocationAreaResponse{}, error
	}
	defer response.Body.Close();
	data, err := io.ReadAll(response.Body);
	if err != nil{
		return LocationAreaResponse{}, error
	}
	locationData := LocationAreaResponse{};
	err = json.Unmarshal(data, &locationData);

	if err != nil{
			return LocationAreaResponse{}, error
	}
	cfgCache := *cache;
	cfgCache.Add(url, data);
	return locationData, nil
}

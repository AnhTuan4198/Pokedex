package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(pageURl *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURl != nil {
		url = *pageURl
	}
	fmt.Println("URL:", url)
	// create request
	req, error := http.NewRequest("GET",url, nil);
	if error != nil {
		return LocationAreaResponse{}, error
	}

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
	fmt.Print(err);
	if err != nil{
			return LocationAreaResponse{}, error
	}

	fmt.Println("data:", locationData);
	return locationData, nil
}

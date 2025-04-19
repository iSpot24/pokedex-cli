package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseUrl + "/location-area?offset=0&limit=20"

	if pageURL != nil {
		url = *pageURL
	}

	data, found := client.cache.Get(url)

	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocations{}, err
		}

		client.cache.Add(url, data)
	}

	respLocations := RespLocations{}
	err := json.Unmarshal(data, &respLocations)

	if err != nil {
		return RespLocations{}, err
	}

	return respLocations, nil
}

func (client *Client) GetLocation(area string) (RespLocation, error) {
	url := baseUrl + "/location-area/" + area

	data, found := client.cache.Get(url)

	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocation{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return RespLocation{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocation{}, err
		}

		client.cache.Add(url, data)
	}

	respLocation := RespLocation{}
	err := json.Unmarshal(data, &respLocation)

	if err != nil {
		return RespLocation{}, err
	}

	return respLocation, nil
}

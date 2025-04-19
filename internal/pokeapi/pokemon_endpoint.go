package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseUrl + "/pokemon/" + name

	data, found := client.cache.Get(url)

	if !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon{}, err
		}

		client.cache.Add(url, data)
	}

	respPokemon := RespPokemon{}
	err := json.Unmarshal(data, &respPokemon)

	if err != nil {
		return RespPokemon{}, err
	}

	return respPokemon, nil
}

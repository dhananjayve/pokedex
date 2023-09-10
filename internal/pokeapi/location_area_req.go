package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResp, error) {
	endPoint := "/location-area"
	fullUrl := BaseUrl + endPoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		locationRes := LocationAreaResp{}
		err := json.Unmarshal(data, &locationRes)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationRes, nil
	}
	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationRes := LocationAreaResp{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(fullUrl, data)
	return locationRes, nil
}

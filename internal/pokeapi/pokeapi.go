package pokeapi

import (
	"command-line-arguments/Users/dhananjayverma/workspace/github.com/dhananjayve/pokedex/internal/pokecache/pokecache.go"
	"net/http"
	"time"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient *http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

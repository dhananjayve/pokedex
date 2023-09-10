package main

import "github.com/dhananjayve/pokedex/internal/pokeapi"

type Config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevAreaLocationUrl *string
}

func main() {
	cfg := Config{
		pokeapiClient: pokeapi.NewClient(),
	}
	StartRepl(&cfg)
	// pokeClient := pokeapi.NewClient()

	// res, err := pokeClient.ListLocationAreas()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)
}

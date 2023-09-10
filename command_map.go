package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *Config) error {
	//pokeClient := pokeapi.NewClient()

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	//fmt.Println(res)
	fmt.Println("list of locations are : ")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = res.Next
	cfg.prevAreaLocationUrl = res.Previous
	return nil
}

func callbackMapBack(cfg *Config) error {
	//pokeClient := pokeapi.NewClient()
	if cfg.prevAreaLocationUrl == nil {
		return errors.New("You are on a first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevAreaLocationUrl)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	//fmt.Println(res)
	fmt.Println("list of locations are : ")
	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	cfg.nextLocationAreaUrl = res.Next
	cfg.prevAreaLocationUrl = res.Previous
	return nil
}

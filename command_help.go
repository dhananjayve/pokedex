package main

import "fmt"

func callbackHelp(cfg *Config) error {
	fmt.Println("Welcome ")
	fmt.Println("available command are: ")
	ac := getCommand()
	for _, c := range ac {
		fmt.Printf(" - %s : %s \n", c.name, c.description)
	}
	// fmt.Println("- help")
	// fmt.Println("- exit ")
	fmt.Println("")
	return nil
}

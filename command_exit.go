package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *Config) error {
	fmt.Println("Quiting application...")
	os.Exit(0)
	return nil
}

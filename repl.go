package main

import (
	bufio "bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please enter some text :")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		availableCmd := getCommand()

		cmd, ok := availableCmd[command]
		if !ok {
			fmt.Println("invalid command ", cleaned)
			continue
		}
		err := cmd.callbacks(cfg)
		if err != nil {
			fmt.Println("callback erros ", err)
		}

		// switch command {
		// case "exit":
		// 	os.Exit(0)
		// case "help":
		// 	fmt.Println("Welcome ")
		// 	fmt.Println("available command are: ")
		// 	fmt.Println("- help")
		// 	fmt.Println("- exit ")
		// default:
		// 	fmt.Println("invalid command ", cleaned)
		// 	//fmt.Println("scanned text is", cleaned)
		// }

	}

}

type cliCmd struct {
	name        string
	description string
	callbacks   func(*Config) error
}

func getCommand() map[string]cliCmd {
	return map[string]cliCmd{
		"help": {
			name:        "help",
			description: "get help item",
			callbacks:   callbackHelp,
		},
		"map": {
			name:        "map",
			description: "get next list of location",
			callbacks:   callbackMap,
		},
		"mapback": {
			name:        "mapback",
			description: "get prev list of location",
			callbacks:   callbackMapBack,
		},
		"exit": {
			name:        "exit",
			description: "turn off cli",
			callbacks:   callbackExit,
		},
	}
}

func cleanInput(s string) []string {
	l := strings.ToLower(s)
	w := strings.Fields(l)
	return w
}

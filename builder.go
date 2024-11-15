package main

import (
	"fmt"
	"os"
)

func Mkmod(
	instructions []interface{},
	name string,
	version string,
	id string,
	packageName string,
	mainClass string) {

	for _, element := range instructions {
		element := element.(map[string]interface{})
		command := element["command"]

		switch command {
		case "mkdir":
			mkdir(element["name"].(string))
			break
		default:
			fmt.Printf("\033[0;31mUnknown command: %s\033[0m\n", command)
			break
		}
	}
}

func mkdir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}
}

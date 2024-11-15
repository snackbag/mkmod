package main

import (
	"fmt"
	"os"
	"path"
	"strings"
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
			mkdir(element["name"].(string), name)
			break
		default:
			fmt.Printf("\033[0;31mUnknown command: %s\033[0m\n", command)
			break
		}
	}
}

func mkdir(dir string, name string) {
	err := os.MkdirAll(path.Join(strings.Replace(name, "/", string(os.PathSeparator), -1), dir), 0755)
	if err != nil {
		panic(err)
	}
}

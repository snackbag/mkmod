package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type ModContext struct {
	Platform    string
	Name        string
	Version     string
	ID          string
	PackageName string
	MainClass   string
}

func Mkmod(
	instructions []interface{}, ctx ModContext) {

	for _, element := range instructions {
		element := element.(map[string]interface{})
		command := element["command"]

		switch command {
		case "mkdir":
			mkdir(element["name"].(string), ctx)
			break
		default:
			fmt.Printf("\033[0;31mUnknown command: %s\033[0m\n", command)
			break
		}
	}
}

func mkdir(dir string, ctx ModContext) {
	err := os.MkdirAll(path.Join(strings.Replace(ctx.Name, "/", string(os.PathSeparator), -1), dir), 0755)
	if err != nil {
		panic(err)
	}
}

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
	Executable  string
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

func MkmodString(original string, ctx ModContext) string {
	newVersion := strings.Replace(original, "%mkmod:platform%", ctx.Platform, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:name%", ctx.Name, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:version%", ctx.Version, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:id%", ctx.ID, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:package%", ctx.PackageName, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:package_dir%", strings.Replace(ctx.PackageName, ".", string(os.PathSeparator), -1), -1)
	newVersion = strings.Replace(newVersion, "%mkmod:main%", ctx.MainClass, -1)

	return newVersion
}

func mkdir(dir string, ctx ModContext) {
	err := os.MkdirAll(MkmodString(path.Join(strings.Replace(ctx.Name, "/", string(os.PathSeparator), -1), dir), ctx), 0755)
	if err != nil {
		panic(err)
	}
}

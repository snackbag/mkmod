package main

import (
	"fmt"
	"io"
	"net/http"
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
	SourcesURL  string
	Variables   map[string]string
}

func Mkmod(data map[string]interface{}, result map[string]interface{}, ctx ModContext) {
	if val2, ok := data["variables"]; ok {
		val2 := val2.(map[string]interface{})

		for k, v := range val2 {
			if _, found := ctx.Variables[k]; !found {
				ctx.Variables[k] = fmt.Sprintf("%v", v)
			}
		}
	}

	if val, ok := data["extend"]; ok {
		val := val.(string)
		CreateMod(result, &ctx.Platform, &val, &ctx.Name, ctx.ID, ctx.PackageName, ctx.MainClass, ctx.Executable, &ctx.SourcesURL, ctx.Variables)
		return
	}

	instructions := data["instructions"].([]interface{})
	fmt.Println(ctx.Variables)

	for _, element := range instructions {
		element := element.(map[string]interface{})
		command := element["command"]

		switch command {
		case "mkdir":
			mkdir(element["name"].(string), ctx)
		case "copy":
			copyFiles(element["files"].([]interface{}), element["to"].(string), ctx)
		case "rename":
			rename(element["dir"].(string), element["file"].(string), element["to"].(string), ctx)
		default:
			fmt.Printf("\033[0;31mUnknown command: %s\033[0m\n", command)
		}
	}
}

func MkmodString(original string, ctx ModContext) string {
	newVersion := strings.Replace(original, "%mkmod:platform%", ctx.Platform, -1)

	for k, v := range ctx.Variables {
		placeholder := fmt.Sprintf("%%mkmod:%s%%", k)
		newVersion = strings.ReplaceAll(newVersion, placeholder, v)
	}

	newVersion = strings.Replace(newVersion, "%mkmod:name%", ctx.Name, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:version%", ctx.Version, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:id%", ctx.ID, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:package%", ctx.PackageName, -1)
	newVersion = strings.Replace(newVersion, "%mkmod:package_dir%", strings.Replace(ctx.PackageName, ".", string(os.PathSeparator), -1), -1)
	newVersion = strings.Replace(newVersion, "%mkmod:main%", ctx.MainClass, -1)

	return newVersion
}

func repath(original string) string {
	return strings.Replace(original, "/", string(os.PathSeparator), -1)
}

func interfaceToString(original []interface{}) []string {
	str := make([]string, len(original))
	for i, val := range original {
		str[i] = val.(string)
	}

	return str
}

func mkdir(dir string, ctx ModContext) {
	dir = MkmodString(path.Join(ctx.Name, repath(dir)), ctx)
	fmt.Printf("create dir: %s\n", dir)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
}

func copyFiles(rawFiles []interface{}, to string, ctx ModContext) {
	files := interfaceToString(rawFiles)
	to = MkmodString(to, ctx)

	for _, file := range files {
		base := path.Base(file)
		filePath := path.Join(ctx.Name, repath(to), base)
		out, err := os.Create(filePath)

		if err != nil {
			panic(err)
		}

		fmt.Printf("create file: %s\n", filePath)

		resp, err := http.Get(ctx.SourcesURL + fmt.Sprintf("/%s/%s/files/%s", ctx.Platform, ctx.Version, MkmodString(file, ctx)))
		if err != nil {
			panic(err)
		}

		rawBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
			fmt.Printf("\033[0;31mFailed to get file '%s' from copy instruction -- got code %d | Server: %s\033[0m\n", file, resp.StatusCode, resp.Body)
			return
		}

		body := string(rawBody)

		out.WriteString(MkmodString(body, ctx))

		out.Close()
		resp.Body.Close()
	}
}

func rename(dir string, file string, new string, ctx ModContext) {
	new = MkmodString(new, ctx)
	file = MkmodString(file, ctx)

	basePath := path.Join(ctx.Name, repath(MkmodString(dir, ctx)))
	filePath := path.Join(basePath, file)
	newPath := path.Join(basePath, new)

	fmt.Printf("rename: %s -> %s\n", filePath, new)

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("\033[0;31mFailed to rename file '%s' from rename instruction -- does not exist\033[0m\n", filePath)
		return
	}

	err := os.Rename(filePath, newPath)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

var AppVersion string
var UpdateURL string

func main() {
	AppVersion = "1.1.2"
	UpdateURL = "https://raw.githubusercontent.com/snackbag/mkmod/refs/heads/main/update.json"
	errors := make([]string, 0)

	CheckVersion()

	platform := flag.String("platform", "fabric", "the mod's platform (e.g. fabric)")
	version := flag.String("version", "1.21.3", "the target minecraft version")
	name := flag.String("name", "Template Mod", "the mod's name")
	sources := flag.String("sources", "https://raw.githubusercontent.com/snackbag/mkmod/refs/heads/main", "The place where mkmod gets its data")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("mkmod version %s - use mkmod -help for help\n", AppVersion)
		return
	}

	if len(flag.Args()) != 3 {
		fmt.Println("Invalid Syntax. Use mkmod -platform [platform] -version [mc version] -name [mod name] [mod id] [package] [main class]")
		return
	}

	id := flag.Args()[0]
	packageName := flag.Args()[1]
	mainName := strings.ToUpper(flag.Args()[2][:1]) + flag.Args()[2][1:]

	packageMatch := matchesRegex("^([a-z][a-zA-Z0-9_]*)(\\.[a-z][a-zA-Z0-9_]*)*$", packageName)
	mainMatch := matchesRegex("^[A-Z][a-zA-Z0-9_$]*$", mainName)
	idMatch := matchesRegex("^[a-z0-9_.-]+$", id)

	if !packageMatch {
		errors = append(errors, "Invalid package. Please follow Java's naming conventions.")
	}

	if !mainMatch {
		errors = append(errors, "Invalid main class. Please follow Java's naming conventions.")
	}

	if !idMatch {
		errors = append(errors, "Invalid mod id. May only consist of lowercase characters, numbers, underscore, dot or dash")
	}

	expath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path.Join(expath, *name)); err == nil {
		errors = append(errors, fmt.Sprintf("There is already a file named '%s' in this directory", *name))
	}

	if len(errors) > 0 {
		fmt.Println("\033[1mFailed to generate template due to the following errors:\033[0;31m")
		fmt.Print("* " + strings.Join(errors, "\n* ") + "\n")
		fmt.Println("\033[0m\nYou must resolve these issues before the template can be created")
		return
	}

	fmt.Println("-----    [Template Settings]    -----")
	fmt.Println("\033[0mName:\t\t\t\033[1m", *name)
	fmt.Println("\033[0mPlatform:\t\t\033[1m", *platform)
	fmt.Println("\033[0mMinecraft Version:\t\033[1m", *version)
	fmt.Println("\033[0mMod ID:\t\t\t\033[1m", id)
	fmt.Println("\033[0mJava Package:\t\t\033[1m", packageName)
	fmt.Println("\033[0mJava Main Class:\t\033[1m", mainName)
	fmt.Println("\033[0mCurrent Path:\t\t\033[1m", expath)
	fmt.Println("\033[0m")

	fmt.Print("Confirm creation of new mod template? [Y/n] ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(strings.ToLower(input))

	if strings.ToLower(input) != "y" {
		fmt.Println("Aborted. ")
		return
	}

	startTime := time.Now()

	templateurl := *sources + "/templates.mkmod.json"
	fmt.Printf("Fetching templates at %s\n", templateurl)

	resp, err := http.Get(templateurl)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		fmt.Println("\033[0;31mCould not receive template file\033[0m")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	schema := int(result["schema"].(float64))
	if schema != 1 {
		fmt.Printf("\033[0;31mInvalid schema version: %d, expected 2\033[0m\n", schema)
		return
	}

	CreateMod(result, platform, version, name, id, packageName, mainName, expath, sources, make(map[string]string))
	diff := time.Since(startTime)
	fmt.Printf("\033[0;32mFinished\033[0m generating mod template in %fs\n", diff.Seconds())
}

func CreateMod(result map[string]interface{}, platform *string, version *string, name *string, id string, packageName string, mainName string, expath string, sources *string, variables map[string]string) {
	if _, exists := result[*platform]; !exists {
		fmt.Printf("\033[0;31mTemplate meta is missing platform '%s'. If you believe this should be added, please create a feature request on the issue tracker\033[0m\n", *platform)
		return
	}

	if _, exists := result[*platform].(map[string]interface{})[*version]; !exists {
		fmt.Printf("\033[0;31mmkmod does not have a template for minecraft %s on %s. If you want to add this version, please create a feature suggestion on the issue tracker\033[0m\n", *version, *platform)
		return
	}

	Mkmod(result[*platform].(map[string]interface{})[*version].(map[string]interface{}), result, ModContext{*platform, *name, *version, id, packageName, mainName, expath, *sources, variables})
}

func matchesRegex(regex string, input string) bool {
	match, err := regexp.MatchString(regex, input)
	if err != nil {
		panic(err)
	}

	return match
}

func CheckVersion() {
	resp, err := http.Get(UpdateURL)
	if err != nil {
		fmt.Println("Failed to check for updates")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		fmt.Printf("\033[0;31mCould not receive update file, status code %v\033[0m\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read update file")
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Failed to unmarshal update json")
		return
	}

	latest := result["latest"].(string)
	if latest == AppVersion {
		return
	}

	fmt.Printf("⚠️ \033[1;33mmkmod %s has been released. \033[0;33mYou are still on %s! Please update, or you might encounter unexpected or UNSAFE behavior!\033[0m\n", latest, AppVersion)
	fmt.Println("\033[0;33mInstructions -> https://github.com/snackbag/mkmod\033[0m")
}

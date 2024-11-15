package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"golang.org/x/text/cases"
)

var validPlatforms []string

func main() {
	validPlatforms = []string{"fabric"}
	errors := make([]string, 0)

	platform := flag.String("platform", "fabric", "the mod's platform (e.g. fabric)")
	version := flag.String("version", "1.21.3", "the target minecraft version")
	name := flag.String("name", "Template Mod", "the mod's name")
	sources := flag.String("sources", "https://raw.githubusercontent.com/snackbag/mkmod/refs/heads/main", "The place where mkmod gets its data")

	flag.Parse()

	if len(flag.Args()) != 3 {
		fmt.Println("Invalid Syntax. Use mkmod -platform [platform] -version [mc version] -name [mod name] [mod id] [package] [main class]")
		return
	}

	if !slices.Contains(validPlatforms, *platform) {
		errors = append(errors, "Invalid platform '"+*platform+"'")
	}

	id := flag.Args()[0]
	packageName := flag.Args()[1]
	mainName := cases.Title(language.English, cases.Compact).String(flag.Args()[2])

	packageMatch := matchesRegex("^(?:\\w+|\\w+\\.\\w+)+$", packageName)
	mainMatch := matchesRegex("^[a-zA-Z_$]+$", mainName)
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

	if len(errors) > 0 {
		fmt.Println("\033[1mFailed to generate template due to the following errors:\033[0;31m")
		fmt.Print("* " + strings.Join(errors, "\n* ") + "\n")
		fmt.Println("\033[0m\nYou must resolve these issues before the template can be created")
		return
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	expath := filepath.Dir(ex)

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

	if _, exists := result[*platform]; !exists {
		fmt.Printf("\033[0;31mTemplate meta is missing platform '%s, please report this issue\033[0m\n", *platform)
		return
	}

	if _, exists := result[*platform].(map[string]interface{})[*version]; !exists {
		fmt.Printf("\033[0;31mmkmod does not have a template for minecraft %s on %s. If you want to add this version, please create a feature suggestion on the GitHub repository\033[0m\n", *version, *platform)
		return
	}

	//templateDir := result["properties"].(map[string]interface{})["templateDir"].(string)
}

func matchesRegex(regex string, input string) bool {
	match, err := regexp.MatchString(regex, input)
	if err != nil {
		panic(err)
	}

	return match
}

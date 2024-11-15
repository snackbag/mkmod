package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/text/language"
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

	fmt.Println("Please wait...")
}

func matchesRegex(regex string, input string) bool {
	match, err := regexp.MatchString(regex, input)
	if err != nil {
		panic(err)
	}

	return match
}

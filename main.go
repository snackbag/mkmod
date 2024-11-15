package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	platform := flag.String("platform", "fabric", "the mod's platform (e.g. fabric)")
	version := flag.String("version", "1.21.3", "the target minecraft version")
	name := flag.String("name", "Template Mod", "the mod's name")

	flag.Parse()

	if len(flag.Args()) != 3 {
		fmt.Println("Invalid Syntax. Use mkmod -platform [platform] -version [mc version] -name [mod name] [mod id] [package] [main class]")
		return
	}

	id := flag.Args()[0]
	packageName := flag.Args()[1]
	mainName := flag.Args()[2]

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

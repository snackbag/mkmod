package main

import (
	"flag"
	"fmt"
)

func main() {

	platform := flag.String("platform", "fabric", "the mod's platform (e.g. fabric)")
	version := flag.String("version", "1.21.3", "the target minecraft version")
	name := flag.String("name", "Template Mod", "the mod's name")

	flag.Parse()

	id := flag.Args()[0]
	packageName := flag.Args()[1]
	mainName := flag.Args()[2]

	fmt.Println("platform:", *platform)
	fmt.Println("version:", *version)
	fmt.Println("name:", *name)
	fmt.Println("id:", id)
	fmt.Println("package:", packageName)
	fmt.Println("main:", mainName)
	fmt.Println("tail:", flag.Args())
}

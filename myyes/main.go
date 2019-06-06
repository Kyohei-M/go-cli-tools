package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	help bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
}

func main() {
	flag.Parse()
	if help {
		showUsage()
		return
	}

	var texts []string
	for _, arg := range os.Args[1:] {
		if arg[0] == '$' {
			texts = append(texts, os.Getenv(arg[1:]))
		} else {
			texts = append(texts, arg)
		}
	}

	var output string
	if len(texts) == 0 {
		output = "y"
	} else {
		output = strings.Join(texts, " ")
	}

	for {
		fmt.Println(output)
	}
}

func showUsage() {
	fmt.Println("Usage: yes [STRING]...")
	fmt.Println("  or:  yes OPTION")
	fmt.Println("Repeatedly output a line with all specified STRING(s), or 'y'.")
	fmt.Println("")
	fmt.Println("      --help     display this help and exit")
}

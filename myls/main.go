package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

	var paths []string
	args := flag.Args()
	if len(args) == 0 {
		paths = append(paths, ".")
	} else {
		paths = args
	}

	multi := len(paths) > 1
	for i, path := range paths {
		listFiles(path, multi, i == 0)
	}
}

func showUsage() {
	fmt.Println("Usage: ls [FILE]...")
	fmt.Println("List information about the FILEs (the current directory by default).")
}

func listFiles(path string, multi, first bool) {
	if multi {
		if !first {
			fmt.Println("")
			fmt.Println("")
		}
		fmt.Println(path + ":")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		if i > 0 {
			fmt.Print("  ")
		}
		fmt.Print(file.Name())
	}
}

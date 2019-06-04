package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("chmod: missing operand\nTry 'chmod --help' for more information.\n")
		return
	} else if len(args) == 1 {
		fmt.Printf("chmod: missing operand after ‘%s’\nTry 'chmod --help' for more information.\n", args[0])
		return
	}

	add := args[0][0] == '+'
	remove := args[0][0] == '-'
	mode := getMode(args[0], add || remove)

	if add {
		for _, file := range args[1:] {
			addPerm(file, mode)
		}
	} else if remove {
		for _, file := range args[1:] {
			removePerm(file, mode)
		}
	} else {
		for _, file := range args[1:] {
			changePerm(file, mode)
		}
	}
}

func getMode(arg string, allowChar bool) int {
	mode := 0
	if allowChar {
		if strings.Contains(arg, "x") {
			mode++
		}
		if strings.Contains(arg, "w") {
			mode += 1 << 1
		}
		if strings.Contains(arg, "r") {
			mode += 1 << 2
		}
	}

	if mode == 0 {
		val, err := strconv.ParseInt(arg, 8, 32)
		mode = int(val)
		if err != nil {
			log.Fatal(err)
		}
	}
	return mode
}

func changePerm(file string, mode int) {
	err := os.Chmod(file, os.FileMode(mode))
	if err != nil {
		fmt.Println(err)
	}
}

func addPerm(file string, mode int) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	present := info.Mode()
	mode = int(present) | mode
	err = os.Chmod(file, os.FileMode(mode))
	if err != nil {
		fmt.Println(err)
	}
}

func removePerm(file string, mode int) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	present := info.Mode()
	mode = (int(present) | mode) ^ mode
	err = os.Chmod(file, os.FileMode(mode))
	if err != nil {
		fmt.Println(err)
	}
}

func showUsage() {
	fmt.Println("Usage: chmod [OPTION]... MODE[,MODE]... FILE...")
	fmt.Println("  or:  chmod [OPTION]... OCTAL-MODE FILE...")
	fmt.Println("Change the mode of each FILE to MODE.")
	fmt.Println("")
	fmt.Println("      --help     display this help and exit")
}

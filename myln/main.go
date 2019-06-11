package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	help bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
}

func main() {
	flag.Parse()
	if help {
		showUsage()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("ln: missing file operand\nTry 'ln --help' for more information.\n")
		return
	}

	if len(args) == 1 {
		target := args[0]
		name, err := getFileName(target)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Symlink(target, filepath.Join(".", name))
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if len(args) == 2 {
		target := args[0]
		link := args[1]
		err := os.Symlink(target, link)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		dir := args[len(args)-1]
		for _, arg := range args[:len(args)-1] {
			name, err := getFileName(arg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = os.Symlink(arg, filepath.Join(dir, name))
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func getFileName(path string) (string, error) {
	f, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return f.Name(), err
}

func showUsage() {
	fmt.Println("Usage: ln [OPTION]... [-T] TARGET LINK_NAME")
	fmt.Println("  or:  ln [OPTION]... TARGET")
	fmt.Println("  or:  ln [OPTION]... TARGET... DIRECTORY")
	fmt.Println("")
	fmt.Println("Mandatory arguments to long options are mandatory for short options too.")
	fmt.Println("    --help     display this help and exit")
}

package main

import (
	"flag"
	"fmt"
	"os"
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
		fmt.Printf("rmdir: missing operand\nTry 'rmdir --help' for more information.\n")
		return
	}

	for _, arg := range args {
		err := validateFilePath(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = os.Remove(arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func validateFilePath(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return fmt.Errorf("rmdir: failed to remove '%s': Not a directory", path)
	}
	return nil
}

func showUsage() {
	fmt.Println("Usage: rmdir [OPTION]... DIRECTORY...")
	fmt.Println("Remove the DIRECTORY(ies), if they are empty.")
	fmt.Println("")
	fmt.Println("      --help     display this help and exit")
}

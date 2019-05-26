package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help      bool
	recursive bool
	directory bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.BoolVar(&recursive, "recursive", false, "remove recursively")
	flag.BoolVar(&recursive, "r", false, "remove recursively (shorthand)")
	flag.BoolVar(&recursive, "R", false, "remove recursively (shorthand)")
	flag.BoolVar(&directory, "dir", false, "remove directory")
	flag.BoolVar(&directory, "d", false, "remove directory (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	args := flag.Args()
	var err error
	for _, arg := range args {
		err = validateFilePath(arg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = remove(arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func showUsage() {
	fmt.Println("Usage: rm [OPTION]... [FILE]...")
	fmt.Println("Remove (unlink) the FILE(s).")
	fmt.Println("")
	fmt.Println("  -r, -R, --recursive   remove directories and their contents recursively")
	fmt.Println("  -d, --dir             remove empty directories")
	fmt.Println("      --help     display this help and exit")
}

func remove(path string) error {
	if recursive {
		return os.RemoveAll(path)
	}
	return os.Remove(path)
}

func validateFilePath(path string) error {
	if !directory && !recursive {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			return fmt.Errorf("rm: cannot remove '%s': Is a directory", path)
		}
	}
	return nil
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	help      bool
	targetDir string
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.StringVar(&targetDir, "target-directory", "", "target directory")
	flag.StringVar(&targetDir, "t", "", "target directory (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	args := flag.Args()
	count := len(args)
	if len(args) == 0 {
		fmt.Println("mv: missing file operand")
		fmt.Println("Try 'mv --help' for more information.")
	} else if count == 1 && targetDir == "" {
		fmt.Printf("mv: missing destination file operand after '%s'/n", args[0])
		fmt.Println("Try 'mv --help' for more information.")
	}

	if targetDir != "" {
		for _, arg := range args {
			move(arg, targetDir)
		}
	} else if count == 2 {
		move(args[0], args[1])
	} else {
		dst := args[count-1]
		for i := 0; i < count-1; i++ {
			move(args[i], dst)
		}
	}
}

func showUsage() {
	fmt.Println("Usage: mv [OPTION]... [-T] SOURCE DEST")
	fmt.Println("  or:  mv [OPTION]... SOURCE... DIRECTORY")
	fmt.Println("  or:  mv [OPTION]... -t DIRECTORY SOURCE...")
	fmt.Println("Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.")
	fmt.Println("")
	fmt.Println("Mandatory arguments to long options are mandatory for short options too.")
	fmt.Println("  -t, --target-directory=DIRECTORY  move all SOURCE arguments into DIRECTORY")
	fmt.Println("      --help     display this help and exit")
}

func move(src, dst string) {
	err := validateSource(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	destination := dst
	if isDir(dst) {
		destination = filepath.Join(dst, filepath.Base(src))
	}

	err = os.Rename(src, destination)
	if err != nil {
		fmt.Println(err)
	}
}

func validateSource(src string) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}
	return nil
}

func isDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return stat.IsDir()
}

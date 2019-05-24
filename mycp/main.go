package main

import (
	"flag"
	"fmt"
	"io"
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
	if count == 0 {
		fmt.Println("cp: missing file operand")
		fmt.Println("Try 'cp --help' for more information.")
		return
	} else if count == 1 && targetDir == "" {
		io.WriteString(os.Stdout, fmt.Sprintln("cp: missing destination file operand after '%s'", args[0]))
		fmt.Println("Try 'cp --help' for more information.")
		return
	}

	if count == 2 && targetDir == "" {
		err := copy(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	} else if targetDir != "" {
		for _, arg := range args {
			copyToDir(arg, targetDir)
		}
	} else {
		dst := args[count-1]
		for i := 0; i < count-1; i++ {
			copyToDir(args[i], dst)
		}
	}
}

func showUsage() {
	fmt.Println("Usage: cp [OPTION]... [-T] SOURCE DEST")
	fmt.Println("  or:  cp [OPTION]... SOURCE... DIRECTORY")
	fmt.Println("  or:  cp [OPTION]... -t DIRECTORY SOURCE...")
	fmt.Println("Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.")
	fmt.Println("")
	fmt.Println("Mandatory arguments to long options are mandatory for short options too.")
	fmt.Println("      --help     display this help and exit")
	fmt.Println("  -t, --target-directory=DIRECTORY  copy all SOURCE arguments into DIRECTORY")
}

func copy(src, dst string) error {
	err := validateSource(src)
	if err != nil {
		return err
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func copyToDir(src, dst string) {
	err := copy(src, filepath.Join(dst, filepath.Base(src)))
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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	help     bool
	noCreate bool
	atime    bool
	mtime    bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.BoolVar(&noCreate, "no-create", false, "do not create any files")
	flag.BoolVar(&noCreate, "c", false, "do not create any files (shorthand)")
	flag.BoolVar(&atime, "a", false, "change only the access time")
	flag.BoolVar(&mtime, "m", false, "change only the modification time")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		touch(arg)
	}
}

func showUsage() {
	fmt.Println("Usage: touch [OPTION]... FILE...")
	fmt.Println("Update the access and modification times of each FILE to the current time.")
	fmt.Println("")
	fmt.Println("A FILE argument that does not exist is created empty, unless -c is supplied.")
	fmt.Println("")
	fmt.Println("  -a                     change only the access time")
	fmt.Println("  -c, --no-create        do not create any files")
	fmt.Println("  -m                     change only the modification time")
	fmt.Println("      --help     display this help and exit")
}

func touch(path string) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		createFile(path)
		return
	}

	changeTimes(path, &fileInfo)
}

func createFile(path string) {
	if noCreate {
		return
	}
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

func changeTimes(path string, fileInfo *os.FileInfo) {
	at := getAtime(fileInfo)
	mt := (*fileInfo).ModTime()
	ct := time.Now().Local()

	if atime && !mtime {
		at = ct
	} else if !atime && mtime {
		mt = ct
	}

	err := os.Chtimes(path, at, mt)
	if err != nil {
		log.Println(err)
	}
}

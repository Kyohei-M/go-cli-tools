package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type modeInt uint32

func (m *modeInt) String() string {
	return fmt.Sprintf("%v", *m)
}

func (m *modeInt) Set(v string) error {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return err
	}

	*m = modeInt(i)
	return nil
}

var (
	help    bool
	mode    modeInt
	parents bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.BoolVar(&parents, "parents", false, "make parent directories as needed")
	flag.BoolVar(&parents, "p", false, "make parent directories as needed (shorthand)")
	flag.Var(&mode, "mode", "set file mode")
	flag.Var(&mode, "m", "set file mode (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	if mode == 0 {
		mode = 0775
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("mkdir: missing operand")
		fmt.Println("Try 'mkdir --help' for more information.")
		return
	}

	for _, arg := range args {
		createDir(arg)
	}
}

func showUsage() {
	fmt.Println("Usage: mkdir [OPTION]... DIRECTORY...")
	fmt.Println("Create the DIRECTORY(ies), if they do not already exist.")
	fmt.Println("")
	fmt.Println("Mandatory arguments to long options are mandatory for short options too.")
	fmt.Println("  -m, --mode=MODE   set file mode (as in chmod), not a=rwx - umask")
	fmt.Println("  -p, --parents     no error if existing, make parent directories as needed")
	fmt.Println("      --help     display this help and exit")
}

func createDir(name string) {
	if _, err := os.Stat(name); !os.IsNotExist(err) {
		fmt.Printf("mkdir: cannot create directory ‘%s’: File exists\n", name)
		return
	}

	if parents {
		os.MkdirAll(name, os.FileMode(mode))
	} else {
		os.Mkdir(name, os.FileMode(mode))
	}
}

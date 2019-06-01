package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
		fmt.Println("kill: usage: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]")
		return
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}

	err = process.Kill()
	if err != nil {
		log.Fatal(err)
	}
}

func showUsage() {
	fmt.Println("kill: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]")
	fmt.Println("    Send a signal to a job.")
	fmt.Println("")
	fmt.Println("    Send the processes identified by PID or JOBSPEC the signal named by")
	fmt.Println("    SIGSPEC or SIGNUM.  If neither SIGSPEC nor SIGNUM is present, then")
	fmt.Println("    SIGTERM is assumed.")
	fmt.Println("")
}

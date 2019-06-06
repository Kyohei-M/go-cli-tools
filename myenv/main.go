package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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
	var subs []string

	for i, arg := range args {
		if strings.Contains(arg, "=") {
			sp := strings.Split(arg, "=")
			if sp[0] != "" {
				fmt.Println("sp=", sp)
				os.Setenv(sp[0], strings.Join(sp[1:], "="))
			}
		} else {
			subs = args[i:]
			break
		}
	}

	fmt.Println("subs=", subs)
	if len(subs) > 0 {
		cmd := exec.Command(subs[0], subs[1:]...)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(stdoutStderr)
	}
}

func getCommandArgs(s []string) []string {
	var texts []string
	for _, arg := range s {
		if arg[0] == '$' {
			texts = append(texts, os.Getenv(arg[1:]))
		} else {
			texts = append(texts, arg)
		}
	}
	return texts
}

func showUsage() {
	fmt.Println("Usage: env [OPTION]... [-] [NAME=VALUE]... [COMMAND [ARG]...]")
	fmt.Println("Set each NAME to VALUE in the environment and run COMMAND.")
	fmt.Println("")
	fmt.Println("Mandatory arguments to long options are mandatory for short options too.")
	fmt.Println("      --help     display this help and exit")
}

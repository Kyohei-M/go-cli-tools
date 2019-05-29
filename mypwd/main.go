package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	help  bool
	flagL bool
	flagW bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.BoolVar(&flagL, "L", false, "print the value of $PWD")
	flag.BoolVar(&flagW, "W", false, "print the Win32 value of the physical directory")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	if runtime.GOOS != "windows" || flagL || !flagW {
		pwd := os.Getenv("PWD")
		if pwd == "" {
			pwd = "/"
		}
		fmt.Println(pwd)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
	}
}

func showUsage() {
	fmt.Println("pwd: pwd [-LPW]")
	fmt.Println("    Print the name of the current working directory.")
	fmt.Println("")
	fmt.Println("    Options:")
	fmt.Println("      -L        print the value of $PWD if it names the current working directory")
	if runtime.GOOS == "windows" {
		fmt.Println("      -W        print the Win32 value of the physical directory")
	}
}

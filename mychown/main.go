package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
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
		fmt.Printf("chown: missing operand\nTry 'chown --help' for more information.\n")
		return
	} else if len(args) == 1 {
		fmt.Printf("chown: missing operand after ‘%s’\nTry 'chown --help' for more information.\n", args[0])
		return
	}

	uid, gid := -1, -1
	sp := strings.Split(args[0], ":")
	if sp[0] != "" {
		user, err := user.Lookup(sp[0])
		if err != nil {
			log.Fatal(err)
		}
		uid, err = strconv.Atoi(user.Uid)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(sp) > 1 && sp[1] != "" {
		group, err := user.LookupGroup(sp[1])
		if err != nil {
			log.Fatal(err)
		}
		gid, err = strconv.Atoi(group.Gid)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range args[1:] {
		err := os.Chown(file, uid, gid)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func showUsage() {
	fmt.Println("Usage: chown [OPTION]... [OWNER][:[GROUP]] FILE...")
	fmt.Println("Change the owner and/or group of each FILE to OWNER and/or GROUP.")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("      --help     display this help and exit")
}

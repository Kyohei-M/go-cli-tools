package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var texts []string
	for _, arg := range os.Args[1:] {
		if arg[0] == '$' {
			texts = append(texts, os.Getenv(arg[1:]))
		} else {
			texts = append(texts, arg)
		}
	}

	fmt.Println(strings.Join(texts, " "))
}

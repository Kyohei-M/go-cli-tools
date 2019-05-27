package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	help   bool
	number bool
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.BoolVar(&number, "number", false, "output with line number")
	flag.BoolVar(&number, "n", false, "output with line number (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	args := flag.Args()
	if len(args) < 1 || args[0] == "-" {
		catFileFromStdin()
		return
	}

	for _, arg := range args {
		catFile(arg)
	}
}

func showUsage() {
	fmt.Println("Usage: cat [OPTION]... [FILE]...")
	fmt.Println("Concatenate FILE(s) to standard output.")
	fmt.Println("")
	fmt.Println("With no FILE, or when FILE is -, read standard input.")
	fmt.Println("")
	fmt.Println("  -n, --number             number all output lines")
	fmt.Println("      --help     display this help and exit")
}

func catFile(path string) {
	fp, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	scanTexts(bufio.NewScanner(fp))
}

func catFileFromStdin() {
	scanTexts(bufio.NewScanner(os.Stdin))
}

func scanTexts(scanner *bufio.Scanner) {
	num := 0
	for scanner.Scan() {
		if number {
			num++
			fmt.Printf("%6d  %s\n", num, scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

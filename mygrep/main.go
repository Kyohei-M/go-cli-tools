package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	help         bool
	lineNumber   bool
	withFileName bool
	noFileName   bool
	maxCount     int
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&lineNumber, "line-number", false, "show line number")
	flag.BoolVar(&lineNumber, "n", false, "show line number (shorthand)")
	flag.BoolVar(&withFileName, "with-filename", false, "print file nane")
	flag.BoolVar(&withFileName, "H", false, "print file name (shorthand)")
	flag.BoolVar(&noFileName, "no-filename", false, "suppress file name")
	flag.BoolVar(&noFileName, "h", false, "suppress file name (shorthand)")
	flag.IntVar(&maxCount, "max-count", 0, "stop after NUM matches")
	flag.IntVar(&maxCount, "m", 0, "stop after NUM matches (shorthand)")
}

func main() {
	flag.Parse()
	if help {
		showUsage()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: grep [OPTION]... PATTERN [FILE]...\nTry 'grep --help' for more information.\n")
	}

	re := regexp.MustCompile(args[0])
	if !withFileName {
		withFileName = len(args) > 2
	}

	if len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		ln := 1
		matchCount := 0
		for scanner.Scan() {
			if maxCount > 0 && maxCount <= matchCount {
				break
			}
			if re.MatchString(scanner.Text()) {
				printResult("", scanner.Text(), ln)
				matchCount++
			}
			ln++
		}
	} else {
		for _, arg := range args[1:] {
			fp, err := os.Open(arg)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer fp.Close()

			scanner := bufio.NewScanner(fp)
			ln := 1
			matchCount := 0
			for scanner.Scan() {
				if maxCount > 0 && maxCount <= matchCount {
					break
				}
				if re.MatchString(scanner.Text()) {
					printResult(arg, scanner.Text(), ln)
					matchCount++
				}
				ln++
			}
		}
	}
}

func printResult(fileName, text string, ln int) {
	var prefix string
	if withFileName && fileName != "" {
		prefix = fmt.Sprint(fileName, ":")
	}
	if noFileName {
		prefix = ""
	}
	if lineNumber {
		prefix = fmt.Sprint(prefix, ln, ":")
	}
	fmt.Printf("%s%s\n", prefix, text)
}

func showUsage() {
	fmt.Println("Usage: grep [OPTION]... PATTERN [FILE]...")
	fmt.Println("Search for PATTERN in each FILE or standard input.")
	fmt.Println("Example: grep -i 'hello world' menu.h main.c")
	fmt.Println("")
	fmt.Println("Miscellaneous:")
	fmt.Println("      --help                display this help text and exit")
	fmt.Println("")
	fmt.Println("Output control:")
	fmt.Println("  -m, --max-count=NUM       stop after NUM matches")
	fmt.Println("  -n, --line-number         print line number with output lines")
	fmt.Println("  -H, --with-filename       print the file name for each match")
	fmt.Println("  -h, --no-filename         suppress the file name prefix on output")
}

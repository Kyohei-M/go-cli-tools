package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
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
		fmt.Printf("sleep: missing operand\nTry 'sleep --help' for more information.")
	}

	time.Sleep(parseDuration(args[0]))
}

func parseDuration(str string) time.Duration {
	durationRegex := regexp.MustCompile(`(?P<days>\d+d)?(?P<hours>\d+h)?(?P<minutes>\d+m)?(?P<seconds>\d+s?$)?`)
	matches := durationRegex.FindStringSubmatch(str)

	if matches[1] == "" && matches[2] == "" && matches[3] == "" && matches[4] == "" {
		log.Fatal(fmt.Printf("sleep: invalid time interval ‘%s’\nTry 'sleep --help' for more information.\n", str))
	}

	days := parseInt64(matches[1])
	hours := parseInt64(matches[2])
	minutes := parseInt64(matches[3])
	seconds := parseSecond(matches[4])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	return time.Duration(days*24*hour + hours*hour + minutes*minute + seconds*second)
}

func parseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}

func parseSecond(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	s := value
	if value[len(value)-1] == 's' {
		s = value[:len(value)-1]
	}
	parsed, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int64(parsed)
}

func showUsage() {
	fmt.Println("Usage: sleep NUMBER[SUFFIX]...")
	fmt.Println("  or:  sleep OPTION")
	fmt.Println("Pause for NUMBER seconds.  SUFFIX may be 's' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days.")
	fmt.Println("")
	fmt.Println("      --help     display this help and exit")
}

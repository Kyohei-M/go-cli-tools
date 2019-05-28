package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type sslice []string

func (s *sslice) String() string {
	return fmt.Sprintf("%v", headers)
}

func (s *sslice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

var (
	help     bool
	method   string
	postData string
	headers  sslice
)

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.StringVar(&method, "X", "GET", "requset command")
	flag.StringVar(&postData, "d", "", "HTTP POST data")
	flag.Var(&headers, "header", "custom header")
	flag.Var(&headers, "H", "custom header (shorthand)")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("curl: try 'curl --help' or 'curl --manual' for more information")
		return
	}

	url := args[0]
	req, err := http.NewRequest(method, url, strings.NewReader(postData))
	if err != nil {
		log.Fatal(err)
	}

	if len(headers) > 0 {
		for _, header := range headers {
			slice := strings.Split(header, ":")
			if len(slice) > 1 {
				req.Header.Add(strings.TrimSpace(slice[0]), strings.TrimSpace(slice[1]))
			}
		}
	}

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArray))
}

func showUsage() {
	fmt.Println("Usage: curl [options...] <url>")
	fmt.Println("")
	fmt.Println(" -d, --data <data>   HTTP POST data")
	fmt.Println(" -H, --header <header/@file> Pass custom header(s) to server")
	fmt.Println(" -h, --help          This help text")
	fmt.Println(" -X, --request <command> Specify request command to use")
}

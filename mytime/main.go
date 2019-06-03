package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]
	var (
		elapsed      time.Duration
		stdoutStderr []byte
	)

	if len(args) > 0 {
		cmd := exec.Command(args[0], args[1:]...)
		start := time.Now()
		var err error
		stdoutStderr, err = cmd.CombinedOutput()
		elapsed = time.Since(start)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("%s\n", stdoutStderr)

	fmt.Printf("real %4.0fm%.3fs\n", elapsed.Round(time.Minute).Minutes(), elapsed.Round(time.Millisecond).Seconds())
}

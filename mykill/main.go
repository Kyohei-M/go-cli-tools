package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"syscall"
)

var (
	help bool
	s    string
	n    int
	l    bool
)

var sigList = []string{
	"SIGHUP",
	"SIGINT",
	"SIGQUIT",
	"SIGILL",
	"SIGTRAP",
	"SIGABRT",
	"SIGEMT",
	"SIGFPE",
	"SIGKILL",
	"SIGBUS",
	"SIGSEGV",
	"SIGSYS",
	"SIGPIPE",
	"SIGALRM",
	"SIGTERM",
	"SIGURG",
	"SIGSTOP",
	"SIGTSTP",
	"SIGCONT",
	"SIGCHLD",
	"SIGTTIN",
	"SIGTTOU",
	"SIGIO",
	"SIGXCPU",
	"SIGXFSZ",
	"SIGVTALRM",
	"SIGPROF",
	"SIGWINCH",
	"SIGPWR",
	"SIGUSR1",
	"SIGUSR2",
	"SIGRTMIN",
	"SIGRTMIN+1",
	"SIGRTMIN+2",
	"SIGRTMIN+3",
	"SIGRTMIN+4",
	"SIGRTMIN+5",
	"SIGRTMIN+6",
	"SIGRTMIN+7",
	"SIGRTMIN+8",
	"SIGRTMIN+9",
	"SIGRTMIN+10",
	"SIGRTMIN+11",
	"SIGRTMIN+12",
	"SIGRTMIN+13",
	"SIGRTMIN+14",
	"SIGRTMIN+15",
	"SIGRTMIN+16",
	"SIGRTMAX-15",
	"SIGRTMAX-14",
	"SIGRTMAX-13",
	"SIGRTMAX-12",
	"SIGRTMAX-11",
	"SIGRTMAX-10",
	"SIGRTMAX-9",
	"SIGRTMAX-8",
	"SIGRTMAX-7",
	"SIGRTMAX-6",
	"SIGRTMAX-5",
	"SIGRTMAX-4",
	"SIGRTMAX-3",
	"SIGRTMAX-2",
	"SIGRTMAX-1",
	"SIGRTMAX",
}

func init() {
	flag.BoolVar(&help, "help", false, "show usage")
	flag.BoolVar(&help, "h", false, "show usage (shorthand)")
	flag.StringVar(&s, "s", "", "signal name")
	flag.IntVar(&n, "n", 0, "signal number")
	flag.BoolVar(&l, "l", false, "list the signal names")
}

func main() {
	flag.Parse()

	if help {
		showUsage()
		return
	}

	if l {
		showSigList()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("kill: usage: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]")
		return
	}

	if n == 0 && s == "" {
		n = 15
	} else if n == 0 && s != "" {
		n = findSigNumber(s)
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "windows" {
		err = process.Kill()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = process.Signal(syscall.Signal(n))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func findSigNumber(sigName string) int {
	for i, sig := range sigList {
		if sig == sigName {
			return i
		}
	}
	log.Fatal(fmt.Errorf("bash: kill: %s: invalid signal specification", sigName))
	return 0
}

func showUsage() {
	fmt.Println("kill: kill [-s sigspec | -n signum | -sigspec] pid | jobspec ... or kill -l [sigspec]")
	fmt.Println("    Send a signal to a job.")
	fmt.Println("")
	fmt.Println("    Send the processes identified by PID or JOBSPEC the signal named by")
	fmt.Println("    SIGSPEC or SIGNUM.  If neither SIGSPEC nor SIGNUM is present, then")
	fmt.Println("    SIGTERM is assumed.")
	fmt.Println("")
	fmt.Println("    Options:")
	fmt.Println("      -s sig    SIG is a signal name")
	fmt.Println("      -n sig    SIG is a signal number")
	fmt.Println("      -l        list the signal names; if arguments follow `-l' they are assumed to be signal numbers for which names should be listed")
}

func showSigList() {
	num := 1
	for _, sig := range sigList {
		fmt.Printf("%2d) %-12s", num, sig)
		num++
	}
}

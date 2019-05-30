// +build linux,!windows

package main

import (
	"os"
	"syscall"
	"time"
)

func getAtime(fileInfo *os.FileInfo) time.Time {
	stat := (*fileInfo).Sys().(*syscall.Stat_t)
	return stat.Atim
}

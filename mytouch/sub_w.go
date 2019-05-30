// +build windows,!linux

package main

import (
	"os"
	"syscall"
	"time"
)

func getAtime(fileInfo *os.FileInfo) time.Time {
	stat := (*fileInfo).Sys().(*syscall.Win32FileAttributeData)
	return time.Unix(0, stat.LastAccessTime.Nanoseconds())
}

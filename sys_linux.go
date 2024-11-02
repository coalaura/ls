//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func isOwnedBy(abs, userUid string) bool {
	info, err := os.Stat(abs)
	if err != nil {
		return false
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return false
	}

	return userUid == fmt.Sprintf("%d", stat.Uid)
}

func absolute(path string) (string, error) {
	return filepath.Abs(path)
}

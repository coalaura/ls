//go:build windows
// +build windows

package main

import (
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

func isOwnedBy(abs, userUid string) bool {
	sd, err := windows.GetNamedSecurityInfo(abs, windows.SE_FILE_OBJECT, windows.OWNER_SECURITY_INFORMATION)
	if err != nil {
		return false
	}

	owner, _, err := sd.Owner()
	if err != nil {
		return false
	}

	return owner.String() == userUid
}

func absolute(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	abs = strings.ReplaceAll(abs, "/", "\\")

	return abs, nil
}

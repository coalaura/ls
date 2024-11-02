package main

import (
	"os"

	"golang.org/x/term"
)

func getTerminalColumns() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 120
	}

	return width
}

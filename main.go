package main

import (
	"os"
	"strings"

	"github.com/coalaura/logger"
)

var log = logger.New()

func main() {
	path := strings.Join(os.Args[1:], " ")
	if path == "" {
		path = "."
	}

	table, err := NewFileTable(path)
	if err != nil {
		log.FatalF("Failed to initialize file table: %v", err)

		return
	}

	if err := table.Process(); err != nil {
		log.FatalF("Failed to read files: %v", err)

		return
	}

	table.Print()
}

package main

import (
	"fmt"
)

type FileInfo interface {
	Name() string
	IsDir() bool
}

type File struct {
	Name      string
	IsDir     bool
	Length    int
	OwnedByUs bool
}

func NewFile(info FileInfo, abs, userUid string) File {
	name := []rune(info.Name())

	if len(name) > 20 {
		end := name[len(name)-8:]

		name = append(name[:8], []rune("...")...)
		name = append(name, end...)
	}

	return File{
		Name:      string(name),
		IsDir:     info.IsDir(),
		Length:    len(name) + 2,
		OwnedByUs: isOwnedBy(abs, userUid),
	}
}

func (f File) String(length int) string {
	var color string

	if !f.OwnedByUs {
		color = "\x1b[0;90m"
	} else if f.IsDir {
		color = "\x1b[1;32m"
	} else {
		color = "\x1b[0;37m"
	}

	return fmt.Sprintf("%s%-*s  \x1b[0m", color, length-2, f.Name)
}

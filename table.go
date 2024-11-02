package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strings"
)

type FileTable struct {
	path string
	user string

	Rows []string
}

func NewFileTable(path string) (*FileTable, error) {
	abs, err := absolute(path)
	if err != nil {
		return nil, err
	}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &FileTable{
		path: abs,
		user: usr.Uid,
	}, nil
}

func (t *FileTable) Process() error {
	var (
		total int
		size  int
		file  File
		files []File
	)

	if strings.Contains(t.path, "*") {
		glob, err := filepath.Glob(t.path)
		if err != nil {
			return err
		}

		for _, path := range glob {
			info, err := os.Stat(path)
			if err != nil {
				return err
			}

			file = NewFile(info, path, t.user)

			files = append(files, file)
			total++

			if size < file.Length {
				size = file.Length
			}
		}
	} else {
		entries, err := os.ReadDir(t.path)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			abs := filepath.Join(t.path, entry.Name())

			file = NewFile(entry, abs, t.user)

			files = append(files, file)
			total++

			if size < file.Length {
				size = file.Length
			}
		}
	}

	if total == 0 {
		text := "No files found."

		t.Rows = []string{
			header(len(text), t.path),
			fmt.Sprintf("\x1b[0;31m%s\x1b[0m", text),
		}

		return nil
	}

	sort.Slice(files, func(i, j int) bool {
		if files[i].IsDir != files[j].IsDir {
			return files[i].IsDir
		}

		return files[i].Name < files[j].Name
	})

	width := getTerminalColumns()

	columns := width / size
	if total < columns {
		columns = total
	}

	t.Rows = []string{
		header(columns*size, t.path),
	}

	var row int

	for i, file := range files {
		if i%columns == 0 {
			t.Rows = append(t.Rows, file.String(size))

			row++
		} else {
			t.Rows[row] += file.String(size)
		}
	}

	return nil
}

func (t *FileTable) Print() {
	for _, row := range t.Rows {
		fmt.Println(row)
	}
}

func header(width int, path string) string {
	length := len(path)

	avail := (width - (length + 2)) / 2
	if avail < 1 {
		avail = 1
	}

	border := strings.Repeat("-", avail)

	return fmt.Sprintf("\x1b[0;90m%s \x1b[0;37m%s \x1b[0;90m%s\x1b[0m", border, path, border)
}

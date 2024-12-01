package helper

import (
	"io"
	"os"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic("failed to open file")
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		panic("failed to read file")
	}

	return string(contents)
}

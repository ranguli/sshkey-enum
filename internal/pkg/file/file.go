package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Exists(filename string) bool {
	exists := true

	_, err := os.OpenFile(filename, os.O_RDONLY, 0o444)
	if errors.Is(err, os.ErrNotExist) {
		exists = false
	}

	return exists
}

func FileErrorMessage(file string) string {
	return fmt.Sprintf("Could not open file \"%s\".\n", file)
}

func ReadLines(path string) ([]string, error) {
	// https://stackoverflow.com/a/18479916
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

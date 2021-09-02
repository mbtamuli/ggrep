package grep

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// ListFiles lists the files in a directory.
func ListFiles(dir string) ([]string, error) {
	filePaths := []string{}
	err := filepath.WalkDir(dir,
		func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

// Search reads slice of strings, finds the pattern
// returns a slice of output if match is found.
func Search(fileContents []string, pattern string) map[int]string {
	output := make(map[int]string)
	for index, line := range fileContents {
		if strings.Contains(line, pattern) {
			output[index] = fileContents[index]
		}
	}
	return output
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
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

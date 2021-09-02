package grep

import (
	"io/fs"
	"path/filepath"
)

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

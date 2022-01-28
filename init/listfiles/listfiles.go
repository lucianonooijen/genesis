package listfiles

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func shouldSkipFile(path string) bool {
	if strings.HasPrefix(path, ".git") {
		return true
	}

	if strings.Contains(path, "node_modules") {
		return true
	}

	return false
}

func RecursivelyFromWorkingDirectory() ([]string, error) {
	var paths []string

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if shouldSkipFile(path) {
			return nil
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	return paths, err
}

package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func expandAbsPath(path string) (string, error) {
	path = os.ExpandEnv(path)

	if len(path) != 0 && !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		path = filepath.Join(wd, path)

	}
	return path, nil
}

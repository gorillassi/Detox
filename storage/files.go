// storage/files.go
package storage

import (
	"os"
	"path/filepath"
)

func ResolvePath(path string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, path[2:])
}

func EnsureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0700)
	}
	return nil
}

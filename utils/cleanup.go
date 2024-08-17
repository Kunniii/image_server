package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

const tempDir = "images"
const maxAge = 24 * time.Hour

// CleanUpOldFiles removes files older than maxAge from the temp directory
func CleanUpOldFiles() error {
	now := time.Now()
	return filepath.WalkDir(tempDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileInfo, err := d.Info()
			if err != nil {
				return err
			}
			fileAge := now.Sub(fileInfo.ModTime())
			if fileAge > maxAge {
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

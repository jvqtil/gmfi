package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func shortHome(path string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	if strings.HasPrefix(path, homeDir) {
		return "~" + path[len(homeDir):]
	}
	return path
}

func fileCmd(path string) (string, error) {
	cmd := exec.Command("file", path)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(output))

	// Remove "<filename>:" prefix
	if idx := strings.Index(result, ":"); idx != -1 {
		result = strings.TrimSpace(result[idx+1:])
	}

	return result, nil
}

func dirSize(root string) (int64, error) {
	var total int64
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return nil
			}
			total += info.Size()
		}
		return nil
	})
	return total, err
}

package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

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

func wrap(arg string) string {
	return "[ " + arg + " ]"
}

func printer(label string, value string, colorFunc func(a ...interface{}) string) {
	if value != "" {
		fmt.Printf("%s %s\n", colorFunc(fmt.Sprintf("%-14s", label)), wrap(value))
	}
}

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

func getSize(size int64) string {
	var fSize, unit string
	switch {
		case size >= 1024*1024*1024*1024: 
		fSize = fmt.Sprintf("%.2f", float64(size)/(1024*1024*1024*1024))
		unit = "TB"
		case size >= 1024*1024*1024: 
		fSize = fmt.Sprintf("%.2f", float64(size)/(1024*1024*1024))
		unit = "GB"
		case size >= 1024*1024: 
		fSize = fmt.Sprintf("%.2f", float64(size)/(1024*1024))
		unit = "MB"
		case size >= 1024: 
		fSize = fmt.Sprintf("%.2f", float64(size)/(1024))
		unit = "KB"
		default: 
		fSize = fmt.Sprintf("%d", size)
		unit = "bytes"
	}

	file := fmt.Sprint(fSize + " " + unit)

	return file
}

func toStr(arg any) string {
	return fmt.Sprint(arg)
}

func wrap(arg string) string {
	return "[ " + arg + " ]"
}

func printer(label string, value string, colorFunc func(a ...interface{}) string) {
	if value != "" {
		fmt.Printf("%s %s\n", colorFunc(fmt.Sprintf("%-14s", label)), wrap(value))
	}
}

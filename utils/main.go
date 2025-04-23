package utils

import "fmt"

func GetSize(size int64) string {
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

func Tostr(arg any) string {
	return fmt.Sprint(arg)
}

func Wrap(arg string) string {
	return "[ " + arg + " ]"
}

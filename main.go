package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	pink   = color.New(color.FgMagenta).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func main() {
	usageHelp := "\n" + "Usage: " + green("gmfi") + blue(" <filename>")

	if len(os.Args) < 2 {
		fmt.Println(usageHelp)
		return
	}

	filename := os.Args[1]
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(red("\nNot found"), wrap(filename))
		} else {
			fmt.Println(red("\nError"), wrap(err.Error()))
		}
		return
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("\nError", wrap(err.Error()))
		return
	}

	var sizeBytes int64
	if info.IsDir() {
		sizeBytes, err = dirSize(filename)
		if err != nil {
			sizeBytes = info.Size()
		}
	} else {
		sizeBytes = info.Size()
	}
	fileSize := getSize(sizeBytes)

	// Print all output
	fmt.Println()
	printer("Object Name", info.Name(), red)
	printer("Object Size", fileSize, green)
	printer("Permissions", toStr(info.Mode()), blue)
	printer("Is Directory", toStr(info.IsDir()), yellow)
	printer("Absolute Path", absPath, pink)
	printer("Last Modified", info.ModTime().Format(time.RFC1123), cyan)
}

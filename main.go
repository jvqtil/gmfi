package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dustin/go-humanize"
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
	if len(os.Args) < 2 {
		fmt.Printf("\nusage: %s %s\n", green("gmfi"), blue("<filename>"))
		fmt.Printf(yellow("github.com/jvqtil/gmfi\n\n"))
		return
	}

	filename := os.Args[1]
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf(red("\nnot found %s :(\n\n"), filename)
		} else {
			fmt.Printf(red("\nerror! %s\n\n"), err.Error)
		}
		return
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf(red("\nerror! %s\n\n"), err.Error)
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
	fileSize := humanize.Bytes(uint64(sizeBytes))

	fileType, err := fileCmd(filename)
	if err != nil {
		fmt.Printf(red("\nerror! %s\n\n"), err.Error)
		return
	}

	permString := fmt.Sprintf("%o", info.Mode().Perm())

	fmt.Printf("\n> %s (%s) - %s [%s]\n",
		red(info.Name()),
		green(fileSize),
		yellow(fileType),
		blue(permString),
	)

	fmt.Printf("%s * %s\n\n",
		pink(shortHome(absPath)),
		cyan(info.ModTime().Format("02 Jan 2006 15:04")),
	)
}

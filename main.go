package main

import (
	"fmt"
	"os"

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
		fmt.Printf("\nusage: %s %s %s\n", green("gmfi"), blue("<filename>"), pink("[or more files]"))
		fmt.Printf(yellow("github.com/jvqtil/gmfi\n"))
		return
	}

	if os.Args[1] == "diff" {
		if len(os.Args) != 4 {
			fmt.Printf("\nusage: %s %s %s\n", green("gmfi"), red("diff"), blue("<file1> <file2>"))
			return
		}
		diffFiles(os.Args[2], os.Args[3])
		return
	}

	for _, filename := range os.Args[1:] {
		showInfo(filename)
	}
}

func showInfo(filename string) {
	meta, err := GetFileMeta(filename)
	if err != nil {
		fmt.Printf(red("\nerror reading %s\n"), filename)
		return
	}

	fmt.Printf("\n> %s (%s) - %s [%s]\n",
		red(meta.Name),
		green(meta.Size),
		yellow(meta.Type),
		blue(meta.Perm),
	)
	fmt.Printf("%s * %s\n",
		pink(meta.Path),
		cyan(meta.Mod),
	)
}

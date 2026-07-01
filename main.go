package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jvqtil/view"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s %s %s\n", green("gmfi"), blue("<filename>"), pink("[or more files]"))
		return
	}

	hidden := true

	for _, arg := range os.Args[2:] {
		if arg == "-h" {
			hidden = false
		}
	}

	switch os.Args[1] {
	case "diff":
		if len(os.Args) != 4 {
			fmt.Printf("usage: %s %s %s\n", green("gmfi"), red("diff"), blue("<what> <with what>"))
			return
		}
		diffFiles(os.Args[2], os.Args[3])

	case "view", "see", "echo", "print":
		if len(os.Args) < 3 {
			fmt.Printf("usage: %s %s %s\n", green("gmfi"), red("view"), blue("<filename>"))
			return
		}
		file := os.Args[2]
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf(red("failed to read file %s\n"), file)
			return
		}
		view.Show(string(data))

	case "big", "small":
		dir := "."
		topN := 5

		if len(os.Args) >= 3 {
			if n, err := strconv.Atoi(os.Args[2]); err == nil {
				topN = n
			} else {
				dir = os.Args[2]
			}
		}

		if len(os.Args) >= 4 {
			if n, err := strconv.Atoi(os.Args[3]); err == nil {
				topN = n
			} else {
				dir = os.Args[3]
			}
		}

		filesSort(os.Args[1], dir, topN)

	case "tree":
		dir := "."
		if len(os.Args) >= 3 && os.Args[2] != "-h" {
			dir = os.Args[2]
		}
		treeCommand(dir, hidden)

	case "search", "find", "grep", "rg":
		if len(os.Args) < 3 {
			fmt.Printf("usage: %s %s %s %s\n", green("gmfi"), red("search"), blue("<pattern>"), pink("[path]"))
			return
		}
		pattern := os.Args[2]
		path := "."
		if len(os.Args) >= 4 {
			path = os.Args[3]
		}
		searchIn(pattern, path)

	case "--help", "-h":
		printHelp()
		return

	case "--version", "-v":
		printVersion()
		return

	default:
		for _, file := range os.Args[1:] {
			showInfo(file)
		}
	}
}

func printHelp() {
	fmt.Printf("usage:\n")
	fmt.Printf(" %s %s %s\n", green("gmfi"), blue("<filename>"), pink("[or more files]"))

	fmt.Printf("commands:\n")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "search")), "find files in directory")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "diff")), "compare two files")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "view")), "print file content with bat, less or cat")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "tree")), "display folder structure")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "big")), "show biggest files in a directory")
	fmt.Printf(" %s > %s\n", blue(fmt.Sprintf("%-6s", "small")), "show smallest files in a directory")

	fmt.Printf(" use %s to exclude hidden files from tree\n", blue("-h"))

	fmt.Printf("flags:\n")
	fmt.Printf(" %s | %s\n", pink("-h"), pink("--help"))
	fmt.Printf(" %s | %s\n", pink("-v"), pink("--version"))

	fmt.Printf("for more, see %s\n", yellow("https://github.com/jvqtil/gmfi/"))
}

func printVersion() {
	fmt.Printf("gmfi %s | %s\n", version, build)
}

func showInfo(file string) {
	meta, err := GetFileMeta(file)
	if err != nil {
		fmt.Printf("%v\n", red(err))
		return
	}

	fmt.Printf("> %s (%s) - %s [%s] | %s\n", red(meta.Name), green(meta.Size), yellow(meta.Type), blue(meta.Perm), cyan(meta.Mod))
}

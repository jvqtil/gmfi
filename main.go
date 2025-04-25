package main

import (
	"fmt"
	"os"
	"path/filepath"
	"flag"
	"time"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	pink := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	helpFlag := flag.Bool("help", false, "Show help and exit")
	hFlag := flag.Bool("h", false, "Show help and exit (shorthand)")
	versionFlag := flag.Bool("version", false, "Show version and exit")
	vFlag := flag.Bool("v", false, "Show version and exit (shorthand)")
	flag.Parse()

	usageHelp := "\n" + "Usage: " + green("gmfi") + blue(" <filename>")

	if *helpFlag || *hFlag {
		const helpText = `
Options:
	-h, --help       Show this help message
	-v, --version    Show version information`
		fmt.Println(usageHelp)
		fmt.Println(helpText)
		return
	}
	if *versionFlag || *vFlag {
		const version = "1.3"
		fmt.Println("\nVersion:", green("gmfi"), wrap(version))
		return
	}

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

	fmt.Printf("%s %s\n",
	red(fmt.Sprintf("%-14s", "Object Name ")),
	wrap(info.Name()))

	fmt.Printf("%s %s\n", 
	green(fmt.Sprintf("%-14s", "Object Size ")), 
	wrap(fileSize))

	fmt.Printf("%s %s\n", 
	blue(fmt.Sprintf("%-14s", "Permissions ")), 
	wrap(toStr(info.Mode())))

	fmt.Printf("%s %s\n", 
	yellow(fmt.Sprintf("%-14s", "Is Directory? ")), 
	wrap(toStr(info.IsDir())))

	fmt.Printf("%s %s\n", 
	pink(fmt.Sprintf("%-14s", "Absolute Path ")), 
	wrap(absPath))

	fmt.Printf("%s %s\n", 
	cyan(fmt.Sprintf("%-14s", "Last Modified ")), 
	wrap(info.ModTime().Format(time.RFC1123)))
}


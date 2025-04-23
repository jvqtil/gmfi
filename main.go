package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jvqtil/gmfi/utils"
	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	pink := color.New(color.FgMagenta).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + green("gmfi") + blue(" <filename>"))
		return
	}

	filename := os.Args[1]
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	file := utils.GetSize(info.Size())
	
	// Print all output
	fmt.Println()

	fmt.Printf("%s %s\n",
	red(fmt.Sprintf("%-14s", "Object Name ")),
	utils.Wrap(info.Name()))

	fmt.Printf("%s %s\n", 
	green(fmt.Sprintf("%-14s", "Object Size ")), 
	utils.Wrap(file))

	fmt.Printf("%s %s\n", 
	blue(fmt.Sprintf("%-14s", "Permissions ")), 
	utils.Wrap(utils.Tostr(info.Mode())))

	fmt.Printf("%s %s\n", 
	yellow(fmt.Sprintf("%-14s", "Is Directory? ")), 
	utils.Wrap(utils.Tostr(info.IsDir())))
	
	fmt.Printf("%s %s\n", 
	pink(fmt.Sprintf("%-14s", "Absolute Path ")), 
	utils.Wrap(absPath))

	fmt.Printf("%s %s\n", 
	cyan(fmt.Sprintf("%-14s", "Last Modified ")), 
	utils.Wrap(info.ModTime().Format(time.RFC1123)))
}


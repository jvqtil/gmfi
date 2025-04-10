package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gmfi <filename>")
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

	fmt.Println("File Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Permissions:", info.Mode())
	fmt.Println("Is Directory:", info.IsDir())
	fmt.Println("Absolute Path", absPath)
	fmt.Println("Last Modified:", info.ModTime().Format(time.RFC1123))
}


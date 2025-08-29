package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getExif(file string) {
	if _, err := exec.LookPath("exiftool"); err != nil {
		fmt.Printf("\n%s\n", red("exiftool is not installed — get it from https://exiftool.org/"))
		return
	}

	cmd := exec.Command("exiftool", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("\nimage meta for %s\n", green(file))
	if cmd.Run() != nil {
		fmt.Printf(red("\nfailed to read metadata for %s :(\n"), file)
	}
}

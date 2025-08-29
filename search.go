package main

import (
	"fmt"
	"os"
	"os/exec"
)

func searchIn(pattern, path string) {
	if _, err := exec.LookPath("rg"); err == nil {
		cmd := exec.Command("rg", "--color=always", pattern, path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if cmd.Run() != nil {
			fmt.Printf(red("\nno matches for '%s' in %s :(\n"), pattern, path)
		}
		return
	}

	if _, err := exec.LookPath("grep"); err == nil {
		cmd := exec.Command("grep", "-rnI", "--color=always", pattern, path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if cmd.Run() != nil {
			fmt.Printf(red("\nno matches for '%s' in %s :(\n"), pattern, path)
		}
		return
	}

	fmt.Printf("\n%s\n", red("neither ripgrep nor grep found in $PATH! please install any to use gmfi search"))
}

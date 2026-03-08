package main

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/fatih/color"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	pink   = color.New(color.FgMagenta).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
)

func shortHome(path string) string {
	home, _ := os.UserHomeDir()
	if strings.HasPrefix(path, home) {
		return "~" + path[len(home):]
	}
	return path
}

func fileCmd(path string) (string, error) {
	out, err := exec.Command("file", path).Output()
	if err != nil {
		return "", err
	}
	s := strings.TrimSpace(string(out))
	if idx := strings.Index(s, ":"); idx != -1 {
		s = strings.TrimSpace(s[idx+1:])
	}
	return s, nil
}

func dirSize(root string) (int64, error) {
	var total int64
	var wg sync.WaitGroup
	paths := make(chan string, 100)
	workerCount := getWorkerCount()

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				info, err := os.Stat(path)
				if err == nil {
					atomic.AddInt64(&total, info.Size())
				}
			}
		}()
	}

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			paths <- path
		}
		return nil
	})
	close(paths)
	wg.Wait()

	return total, nil
}

func dirFileCount(root string) int {
	count := 0
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			count++
		}
		return nil
	})
	return count
}

func getWorkerCount() int {
	count := max(runtime.NumCPU()-1, 1)
	return count
}

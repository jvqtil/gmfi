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

	for i := 0; i < workerCount; i++ {
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
	count := runtime.NumCPU() - 1
	if count < 1 {
		count = 1
	}
	return count
}

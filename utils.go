package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dustin/go-humanize"
)

type FileMeta struct {
	Name    string
	Path    string
	Type    string
	Size    string
	Perm    string
	Mod     string
	RawSize int64
}

func GetFileMeta(path string) (*FileMeta, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	abs, _ := filepath.Abs(path)
	var size int64
	var ftype string
	if info.IsDir() {
		size, _ = dirSize(path)
		count := dirFileCount(path)
		ftype = fmt.Sprintf("directory, %d files", count)
	} else {
		size = info.Size()
		ftype, _ = fileCmd(path)
	}

	return &FileMeta{
		Name:    info.Name(),
		Path:    shortHome(abs),
		Type:    ftype,
		Size:    humanize.Bytes(uint64(size)),
		Perm:    fmt.Sprintf("%o", info.Mode().Perm()),
		Mod:     info.ModTime().Format("02 Jan 2006 15:04"),
		RawSize: size,
	}, nil
}

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

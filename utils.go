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

	"github.com/dustin/go-humanize"
)

const defaultChanBuffer = 100

type FileMeta struct {
	Name, Size, Type, Perm, Mod, Path string
}

type FileMetaSimple struct {
	Name string
	Path string
	Type string
	Size int64
}

func GetFileMeta(path string) (*FileMeta, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	abs, _ := filepath.Abs(path)
	var size int64
	if info.IsDir() {
		size, _ = dirSize(rootClean(path))
	} else {
		size = info.Size()
	}
	ftype, _ := fileCmd(path)

	return &FileMeta{
		Name: info.Name(),
		Size: humanize.Bytes(uint64(size)),
		Type: ftype,
		Perm: fmt.Sprintf("%o", info.Mode().Perm()),
		Mod:  info.ModTime().Format("02 Jan 2006 15:04"),
		Path: shortHome(abs),
	}, nil
}

func GetSimpleFileMeta(path string) (*FileMetaSimple, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	ftype, _ := fileCmd(path)
	return &FileMetaSimple{
		Name: info.Name(),
		Path: path,
		Type: ftype,
		Size: info.Size(),
	}, nil
}

func rootClean(path string) string {
	cleaned, _ := filepath.Abs(path)
	return cleaned
}

func showInfos(files []string) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]*FileMeta, 0, len(files))
	paths := make(chan string, len(files))
	workerCount := getWorkerCount()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				meta, err := GetFileMeta(path)
				if err == nil {
					mu.Lock()
					results = append(results, meta)
					mu.Unlock()
				}
			}
		}()
	}

	for _, f := range files {
		paths <- f
	}
	close(paths)
	wg.Wait()

	for _, meta := range results {
		fmt.Printf("\n> %s (%s) - %s [%s]\n", red(meta.Name), green(meta.Size), yellow(meta.Type), blue(meta.Perm))
		fmt.Printf("%s * %s\n", pink(meta.Path), cyan(meta.Mod))
	}
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
	var mu sync.Mutex
	var wg sync.WaitGroup
	paths := make(chan string, defaultChanBuffer)
	workerCount := getWorkerCount()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				info, err := os.Stat(path)
				if err == nil {
					mu.Lock()
					total += info.Size()
					mu.Unlock()
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

func getWorkerCount() int {
	count := runtime.NumCPU() - 1
	if count < 1 {
		count = 1
	}
	return count
}

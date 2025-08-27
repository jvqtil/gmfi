package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func bigFiles(root string, topN int) {
	var files []FileMeta
	var mu sync.Mutex
	var wg sync.WaitGroup
	paths := make(chan string, 100)
	count := 0

	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			count++
		}
		return nil
	})

	bar := progressbar.Default(int64(count), green(fmt.Sprintf("scanning files in %s", shortHome(root))))
	workerCount := getWorkerCount()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				meta, err := GetFileMeta(path)
				if err == nil {
					mu.Lock()
					files = append(files, *meta)
					mu.Unlock()
				}
				bar.Add(1)
			}
		}()
	}

	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			paths <- path
		}
		return nil
	})
	close(paths)
	wg.Wait()

	if len(files) == 0 {
		fmt.Println("No files found")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].RawSize > files[j].RawSize
	})

	if topN > len(files) {
		topN = len(files)
	}
	if root == "." {
		root = "current directory"
	}

	fmt.Printf("\n%d biggest files in %s\n", topN, shortHome(root))
	for i := 0; i < topN; i++ {
		f := files[i]
		fmt.Printf("\n%-10s %s\n", green(f.Size), pink(shortHome(f.Path)))
		fmt.Printf("%s\n", yellow(f.Type))
	}
}

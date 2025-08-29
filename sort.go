package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func filesSort(sortType string, root string, topN int) {
	var files []FileMeta
	var mu sync.Mutex
	var wg sync.WaitGroup
	paths := make(chan string, 100)

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
		fmt.Printf("\n%s\n", "no files found :(")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		if sortType == "big" {
			return files[i].RawSize > files[j].RawSize
		} else if sortType == "small" {
			return files[i].RawSize < files[j].RawSize
		}
		return false
	})

	if topN > len(files) {
		topN = len(files)
	}
	if root == "." {
		root = "current directory"
	}

	title := "files"
	if sortType == "big" {
		title = "biggest files"
	} else if sortType == "small" {
		title = "smallest files"
	}

	fmt.Printf("\n%d %s in %s\n", topN, title, shortHome(root))
	for i := 0; i < topN; i++ {
		f := files[i]
		fmt.Printf("\n%-10s %s\n", green(f.Size), pink(shortHome(f.Path)))
		fmt.Printf("%s\n", yellow(f.Type))
	}
}

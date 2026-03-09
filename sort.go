package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func filesSort(sortBy string, root string, topN int) {
	var files []FileMeta
	var mu sync.Mutex
	var wg sync.WaitGroup
	paths := make(chan string, 100)

	workerCount := getWorkerCount()

	for range workerCount {
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
		fmt.Printf("\n%s\n", "no files found")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		switch sortBy {

		case "big":
			return files[i].RawSize > files[j].RawSize
		case "small":
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
	switch sortBy {
	case "big":
		title = fmt.Sprintf("%s files", bold("biggest"))
	case "small":
		title = fmt.Sprintf("%s files", bold("smallest"))
	}

	fmt.Printf("\n%d %s in %s\n", topN, title, shortHome(root))
	for i := 0; i < topN; i++ {
		f := files[i]
		fmt.Printf("\n%-10s %s %s\n", green(f.Size), pink(shortHome(f.Path)), blue(f.Mod))
		fmt.Printf("%s\n", yellow(f.Type))
	}
}

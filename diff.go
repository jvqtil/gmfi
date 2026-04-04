package main

import (
	"fmt"
	"sync"
)

func diffFiles(f1, f2 string) {
	var meta1, meta2 *FileMeta
	var err1, err2 error
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		meta1, err1 = GetFileMeta(f1)
	}()
	go func() {
		defer wg.Done()
		meta2, err2 = GetFileMeta(f2)
	}()
	wg.Wait()

	if err1 != nil || err2 != nil {
		if err1 != nil {
			fmt.Printf(red("\nerror reading %s\n"), f1)
		}
		if err2 != nil {
			fmt.Printf(red("\nerror reading %s\n"), f2)
		}
		return
	}

	fmt.Printf("\n> Comparing %s with %s\n\n", red(meta1.Name), green(meta2.Name))
	printDiff("Size", meta1.Size, meta2.Size)
	printDiff("Type", meta1.Type, meta2.Type)
	printDiff("Perms", meta1.Perm, meta2.Perm)
}

func printDiff(label, v1, v2 string) bool {
	if v1 == v2 {
		colored := v1
		switch label {
		case "Name":
			colored = red(v1)
		case "Size":
			colored = green(v1)
		case "Type":
			colored = yellow(v1)
		case "Perms":
			colored = blue(v1)
		}
		fmt.Printf("%-10s %s\n", label+":", colored)
		return false
	}
	fmt.Printf("%-10s %s -> %s\n", label+":", red(v1), green(v2))
	return true
}

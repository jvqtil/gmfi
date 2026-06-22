package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func treeCommand(root string, showHidden bool) {
	fmt.Printf("\n%s\n", shortHome(root))
	err := walkTree(root, showHidden, "")
	if err != nil {
		fmt.Printf("%s", red("\nerror reading directory\n"))
	}
}

func walkTree(path string, showHidden bool, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var visibleEntries []os.DirEntry

	for _, entry := range entries {
		if !showHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		visibleEntries = append(visibleEntries, entry)
	}

	for i, entry := range visibleEntries {
		connector := "├── "
		subPrefix := prefix + "│   "
		if i == len(entries)-1 {
			connector = "└── "
			subPrefix = prefix + "    "
		}

		fullPath := filepath.Join(path, entry.Name())
		meta, metaErr := GetFileMeta(fullPath)
		if metaErr != nil {
			fmt.Printf("%s%s%s %s\n", prefix, connector, red(entry.Name()), red("[error]"))
			continue
		}

		displayName := meta.Name
		if entry.IsDir() {
			displayName = blue(displayName + "/")
		} else {
			displayName = red(meta.Name)
		}

		fmt.Printf("%s%s%s (%s, %s)\n",
			prefix, connector,
			displayName,
			green(meta.Size),
			yellow(meta.Type),
		)

		if entry.IsDir() {
			err := walkTree(fullPath, showHidden, subPrefix)
			if err != nil {
				fmt.Println(red(err))
			}
		}
	}

	return nil
}

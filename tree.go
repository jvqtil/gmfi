package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func treeCommand(root string) {
	fmt.Printf("\n%s\n", shortHome(root))
	err := walkTree(root, "")
	if err != nil {
		fmt.Printf(red("\nerror reading directory! %s\n"), err)
	}
}

func walkTree(path string, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
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
			err := walkTree(fullPath, subPrefix)
			if err != nil {
				fmt.Println(red("error:"), err)
			}
		}
	}

	return nil
}

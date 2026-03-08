package main

import (
	"fmt"
	"os"
	"path/filepath"

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
		Mod:     info.ModTime().Format("02 Jan 06 15:04"),
		RawSize: size,
	}, nil
}

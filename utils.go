package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

type FileMeta struct {
	Name, Size, Type, Perm, Mod, Path string
}

func GetFileMeta(path string) (*FileMeta, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	abs, _ := filepath.Abs(path)
	var size int64
	if info.IsDir() {
		size, _ = dirSize(path)
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
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			if info, err := d.Info(); err == nil {
				total += info.Size()
			}
		}
		return nil
	})
	return total, nil
}

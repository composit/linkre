package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Suffix is the file suffix that indicates to the program to create a symlink of that file
const Suffix = ".linkre"

func main() {
	flag.Parse()
	root := flag.Arg(0)

	visit := buildVisit(root)

	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

func buildVisit(root string) filepath.WalkFunc {
	return func(oldRelPath string, f os.FileInfo, err error) error {
		oldPath, err := filepath.Abs(oldRelPath)
		if err != nil {
			return err
		}

		if filepath.Ext(oldPath) == Suffix {
			newPath := strings.TrimPrefix(
				strings.TrimSuffix(oldRelPath, Suffix), root,
			)
			newPath = filepath.Join("/", newPath)

			fmt.Printf("Symlinking %s...\n", filepath.Base(newPath))
			fmt.Printf("old: %s\n", oldPath)
			fmt.Printf("new: %s\n", newPath)

			ok, err := checkNewPath(oldPath, newPath)
			if err != nil {
				return err
			}

			if !ok {
				return nil
			}

			if err := os.Symlink(oldPath, newPath); err != nil {
				return err
			}
			fmt.Println(" created")
		}
		return nil
	}
}

func checkNewPath(oldPath, newPath string) (bool, error) {
	s, err := os.Lstat(newPath)
	if err != nil {
		if os.IsNotExist(err) {
			// file doesn't exist. go ahead with creation
			return true, nil
		}

		return false, err
	}

	// file exists
	if s.Mode()&os.ModeSymlink != 0 {
		li, err := os.Readlink(newPath)
		if err != nil {
			return false, err
		}
		if li == oldPath {
			fmt.Println(" unchanged")
			return false, nil
		}

		return false, fmt.Errorf("%s exists and links to %s", newPath, li)
	}

	return false, fmt.Errorf("%s exists. Aborting", newPath)
}

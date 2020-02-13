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

	visit := func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == Suffix {
			toPath := strings.TrimPrefix(
				strings.TrimSuffix(path, Suffix), root,
			)
			fmt.Printf("Making a symlink of %s\n", path)
			fmt.Printf("to: %s\n", toPath)
			if err := os.Symlink(path, toPath); err != nil {
				return err
			}
		}
		return nil
	}

	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

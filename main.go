package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const SUFFIX = ".linkre"

func main() {
	flag.Parse()
	root := flag.Arg(0)

	visit := func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == SUFFIX {
			fmt.Printf("Making a symlink of %s\n", path)
			toPath := strings.TrimPrefix(
				strings.TrimSuffix(path, SUFFIX), root,
			)
			fmt.Printf("to: %s\n", toPath)
		}
		return nil
	}

	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

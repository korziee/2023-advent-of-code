package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	file, err := fs.ReadFile(os.DirFS(filepath.Dir(".")), "input")
	if err != nil {
		panic(err)
	}

}

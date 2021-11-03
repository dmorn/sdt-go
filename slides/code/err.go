package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenTmp(name string) (*os.File, error) { // HL
	return os.Open(filepath.Join("/tmp", name))
}

func main() {
	f, err := OpenTmp("abc.dat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	f.Close()
}

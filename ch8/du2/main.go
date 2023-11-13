package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var verbose = flag.Bool("verbose", false, "show verbose progress message")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.FileInfo {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	// Create a slice of FileInfo objects.
	var fileInfos []fs.FileInfo

	// Convert each DirEntry object to a FileInfo object and add it to the slice.
	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}

		fileInfos = append(fileInfos, fileInfo)
	}

	return fileInfos
}

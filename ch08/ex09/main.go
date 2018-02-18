package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type folderInfo struct {
	name string
	size int64
}

type folderResult struct {
	nbytes int64
	nfiles int64
}

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {

	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	folder := make(chan folderInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, folder, root)
	}
	go func() {
		n.Wait()
		close(folder)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(1 * time.Second)
	}
	result := map[string]folderResult{}
loop:
	for {
		select {
		case info, ok := <-folder:
			if !ok {
				break loop // fileSizes was closed
			}
			result[info.name] = folderResult{result[info.name].nbytes + info.size, result[info.name].nfiles + 1}
		case <-tick:
			printDiskUsage(result)
		}
	}

	printDiskUsage(result)
}

func printDiskUsage(result map[string]folderResult) {
	for key, val := range result {
		fmt.Printf("%s: %d files  %.1f GB\n", key, val.nfiles, float64(val.nbytes)/1e9)
	}
}

func walkDir(dir string, n *sync.WaitGroup, folder chan<- folderInfo, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, folder, root)
		} else {
			folder <- folderInfo{root, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

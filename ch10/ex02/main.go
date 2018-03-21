package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	_filename := []rune(filename)
	format := string(_filename[len(filename)-3:])

	var err error
	switch format {
	case "tar":
		err = untar(filename)

	case "zip":
		err = unzip(filename)

	default:
		fmt.Println("unknown format")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func unzip(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Println(f.FileInfo().Name())
	}
	return nil
}

func untar(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := tar.NewReader(file)

	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}

		fmt.Println(header.FileInfo().Name())
	}
	return nil
}

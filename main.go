package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type options struct {
	includeHidden bool
	onlyDirs      bool
	fullPath      bool
	du            bool
}

func handleCommandLineArguments() (options, error) {
	includeHidden := true
	onlyDirs := false
	fullPath := false
	du := false

	for _, arg := range os.Args[2:] {
		switch arg {
		case "-a":
			includeHidden = false
		case "-d":
			onlyDirs = true
		case "-f":
			fullPath = true
		case "--du":
			du = true
		}
	}
	return options{includeHidden, onlyDirs, fullPath, du}, nil
}

func printTreeStructure(path string, indent string, options options) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !options.includeHidden && file.Name()[0] == '.' {
			continue
		}
		if file.IsDir() {
			if options.fullPath {
				fmt.Println(filepath.Join(path, file.Name()) + "/")
			} else {
				fmt.Println(indent + "|--" + file.Name() + "/")
			}
			err := printTreeStructure(filepath.Join(path, file.Name()), indent+"|  ", options)
			if err != nil {
				return err
			}
		} else if !options.onlyDirs {
			if options.fullPath {
				fmt.Println(filepath.Join(path, file.Name()))
			} else {
				fmt.Println(indent + "|--" + file.Name())
			}
		}
	}
	return nil
}

func calculateSize(path string, options options) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, file := range files {
		if !options.includeHidden && file.Name()[0] == '.' {
			continue
		}
		if file.IsDir() {
			subdirSize, err := calculateSize(filepath.Join(path, file.Name()), options)
			if err != nil {
				return 0, err
			}
			size += subdirSize
		} else if !options.onlyDirs {
			size += file.Size()
		}
	}
	return size, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tree [directory] [-a] [-d] [-f] [--du]")
		return
	}
	root := os.Args[1]
	options, err := handleCommandLineArguments()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(root + "/")
	err = printTreeStructure(root, "", options)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if options.du {
		size, err := calculateSize(root, options)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("Total Size: %d bytes\n", size)
	}
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// printTree is a recursive function that takes in a directory path, an indentation string,
// flags for full path, only directories, disk usage, and hidden files and returns the total size of the directory
// The function opens the directory and reads the files and subdirectories in it.
// For each file or subdirectory, it checks the flags and indentation and prints the file/subdirectory in a tree format.
// If the file is a directory and the flag for disk usage is true, it calls the function recursively to calculate the total size of the directory
func printTree(path string, indent string, fullPath, onlyDirs, du bool, includeHidden bool) int64 {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer file.Close()

	files, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var size int64
	for _, file := range files {
		// If the flag for full path is not set, the flag for only directories is not set, and the file is hidden and includeHidden flag is not set, skip the file
		if !fullPath && file.Name()[0] == '.' && !onlyDirs && !includeHidden {
			continue
		}
		if file.IsDir() {
			if du {
				size = printTree(filepath.Join(path, file.Name()), indent+"|  ", fullPath, onlyDirs, du, includeHidden)
				fmt.Printf("%s|--%s [%d bytes]\n", indent, file.Name(), size)
			} else if !onlyDirs {
				if fullPath {
					fmt.Println(filepath.Join(path, file.Name()) + "/")
				} else {
					fmt.Println(indent + "|--" + file.Name() + "/")
				}
				printTree(filepath.Join(path, file.Name()), indent+"|  ", fullPath, onlyDirs, du, includeHidden)
			}
		} else if !onlyDirs {
			if fullPath {
				fmt.Println(filepath.Join(path, file.Name()))
			} else {
				fmt.Println(indent + "|--" + file.Name())
			}
			size += file.Size()
		}
	}
	return size
}

func main() {
	// Check if the user provided the directory path as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: tree [directory] [-a] [-d] [-f] [--du]")
		return
	}
	root := os.Args[1]
	// Initialize the flags as
	// Initialize the flags as default values
	includeHidden, onlyDirs, fullPath, du := true, false, false, false
	// Iterate through the command line arguments and set the flags based on the passed options
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
	// Print the root directory name
	fmt.Println(root + "/")
	// Call the printTree function to print the tree structure of the directory
	printTree(root, "", fullPath, onlyDirs, du, includeHidden)
}

### Treeman

Treeman is a Go CLI, which is a basic port of the ```tree``` command in Linux. My main objective for this tool was learning how Go works.

### How it works

The ```main``` function in this script is responsible for parsing the command line arguments passed to the script. It checks if the user provided a directory path as the first argument, and if not, it prints the usage instructions and exits.

The script then defines a set of flags that control how the tree is printed. These flags include whether to include hidden files, whether to only show directories, whether to show the full path of files and directories, and whether to show the disk usage of each directory.

The script then calls the ```printTree``` function, passing in the root directory and the flags as arguments. This function is responsible for recursively traversing the directory tree and printing the tree structure in a nested, indented format.

The ```printTree function``` takes in a directory path, an indentation string, and the flags. It opens the directory and reads the files and subdirectories in it. For each file or subdirectory, it checks the flags and indentation and prints the file/subdirectory in a tree format. If the file is a directory and the flag for disk usage is true, it calls the function recursively to calculate the total size of the directory.

The function uses the ```filepath.Join``` function to join the path of the current directory with the name of the file or subdirectory, and it uses the ```os.Open``` function to open the file or subdirectory. It then uses the ```file.Readdir(-1)``` function to read the contents of the directory, and it uses the ```file.IsDir()``` function to check if the file is a directory.

The function also uses the ```fmt.Println``` and ```fmt.Printf``` functions to print the tree structure in the desired format. The script also uses defer statement to make sure that the file is closed.

### Usage

For windows, download and run the exe. 

For Linux and MacOS, download the binary, and move it to ```usr/local/bin```

Flags:

- directory: the path of the directory you want to print the tree for.
- -a: include hidden files.
- -d: only print subdirectories.
- -f: print the full path for each file and subdirectory.
- --du: report the disk usage for each subdirectory.

For example, running the command ```treeman /home/user/documents -a``` will print the directory structure of the "documents" directory in the home directory and will include hidden files.

### References

https://www.youtube.com/watch?v=XbKSssBftLM&t=225s
https://github.com/evolbioinfo/gotree
https://github.com/a8m/tree
https://chat.openai.com/

### Future improvements

- Benchmark the tool vs native ```tree```
- I failed at setting up CI/CD for building binaries on each deployment, which I hope to learn.
- Add more native features from ```tree```.



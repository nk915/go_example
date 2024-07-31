package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world.")

	path, err := ResolveSymlink("./test/test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(path)
	}
}

func ResolveSymlink(filePath string) (string, error) {
	// Get the file info using Lstat, which doesn't follow symlinks.
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", fmt.Errorf("error getting file info: %v", err)
	}

	// Check if the file is a symbolic link.
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		// Read the symlink to get the original path.
		originalPath, err := os.Readlink(filePath)
		if err != nil {
			return "", fmt.Errorf("error reading symlink: %v", err)
		}
		// Return the original path.
		return originalPath, nil
	}

	return filePath, nil
}

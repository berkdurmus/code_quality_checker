
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: codechecker <path_to_source_code>")
		return
	}
	sourcePath := args[0]
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v
", path, err)
			return err
		}
		if !info.IsDir() {
			fmt.Println("Analyzing file:", path)
			// Add code analysis logic here
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v
", sourcePath, err)
	}
}

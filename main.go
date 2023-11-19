
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: codechecker <path_to_source_code>")
		return
	}
	sourcePath := args[0]

	// Metrics
	totalLines := 0
	totalCommentLines := 0
	totalFunctionCount := 0
	// Additional metrics can be declared here

	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fileMetrics, err := analyzeFile(path)
			if err != nil {
				fmt.Printf("Error analyzing file %s: %v\n", path, err)
				return err
			}
			totalLines += fileMetrics.totalLines
			totalCommentLines += fileMetrics.commentLines
			totalFunctionCount += fileMetrics.functionCount
			// Add additional metric aggregations here
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", sourcePath, err)
		return
	}

	// Print out the metrics
	fmt.Printf("Total Lines: %d\n", totalLines)
	fmt.Printf("Total Comment Lines: %d\n", totalCommentLines)
	fmt.Printf("Total Function Count: %d\n", totalFunctionCount)
	// Print additional metrics here
}

// FileMetrics holds the metrics for a single file
type FileMetrics struct {
	totalLines     int
	commentLines   int
	functionCount  int
	// Additional metrics fields can be added here
}

// analyzeFile analyzes a single Go file and returns its metrics
func analyzeFile(filePath string) (FileMetrics, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return FileMetrics{}, err
	}
	defer file.Close()

	var metrics FileMetrics
	scanner := bufio.NewScanner(file)
	functionRegex := regexp.MustCompile(`^func\s`)

	for scanner.Scan() {
		line := scanner.Text()
		metrics.totalLines++

		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "//") {
			metrics.commentLines++
		}
		if functionRegex.MatchString(trimmedLine) {
			metrics.functionCount++
		}
		// Add additional metric calculations here
	}

	return metrics, scanner.Err()
}

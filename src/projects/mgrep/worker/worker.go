package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

func (r *Results) Add(result Result) {
	r.Inner = append(r.Inner, result)
}

func NewResult(line string, lineNum int, path string) Result {
	return Result{line, lineNum, path}
}

func FindInFile(path string, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	results := Results{make([]Result, 0)}
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			results.Add(NewResult(scanner.Text(), lineNum, path))
		}
		lineNum++
	}

	if len(results.Inner) == 0 {
		return nil
	}

	return &results

}

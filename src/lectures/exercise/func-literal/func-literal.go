//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import "fmt"

func stringIterator(lines []string, fn func(string)) {
	for _, line := range lines {
		fn(line)
	}
}

func runeReport(line string) {
	// Create a map to store the rune counts
	runeCount := make(map[string]int)
	// Iterate over each rune in the line
	for _, r := range line {
		// Determine the rune type
		switch {
		case 'a' <= r && r <= 'z':
			runeCount["letters"]++
		case 'A' <= r && r <= 'Z':
			runeCount["letters"]++
		case '0' <= r && r <= '9':
			runeCount["digits"]++
		case r == ' ':
			runeCount["spaces"]++
		default:
			runeCount["punctuation"]++
		}
	}
	// Print the report
	fmt.Printf("Line: %s\n", line)
	fmt.Printf("Letters: %d\n", runeCount["letters"])
	fmt.Printf("Digits: %d\n", runeCount["digits"])
	fmt.Printf("Spaces: %d\n", runeCount["spaces"])
	fmt.Printf("Punctuation: %d\n", runeCount["punctuation"])
	fmt.Println()
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	stringIterator(lines, runeReport)
}

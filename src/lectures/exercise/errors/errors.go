//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParserError struct {
	msg, input string
}

func (e *TimeParserError) Error() string {
	return fmt.Sprintf("%v: %v", e.msg, e.input)
}

func parseTime(time string) (Time, error) {
	components := strings.Split(time, ":")

	if len(components) != 3 {
		return Time{}, &TimeParserError{"Invalid number of components", time}
	}

	hour, err := strconv.Atoi(components[0])
	if err != nil || hour < 0 || hour > 23 {
		return Time{}, &TimeParserError{fmt.Sprintf("Error Parsing hour %v", err), components[0]}
	}

	minute, err := strconv.Atoi(components[1])
	if err != nil || minute < 0 || minute > 59 {
		return Time{}, &TimeParserError{fmt.Sprintf("Error Parsing hour %v", err), components[1]}
	}

	second, err := strconv.Atoi(components[2])
	if err != nil || second < 0 || second > 59 {
		return Time{}, &TimeParserError{fmt.Sprintf("Error Parsing hour %v", err), components[2]}
	}

	return Time{hour, minute, second}, nil
}

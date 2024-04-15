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

import "testing"

func TestInvalidHour(t *testing.T) {
	// check hour range
	_, err := parseTime("25:00:00")
	if err == nil {
		t.Error("Expected error hour not within valid range, got nil")
	}

	// check string as hour
	_, err = parseTime("a:00:00")
	if err == nil {
		t.Error("Expected error hour not a number, got nil")
	}
}

func TestInvalidMinute(t *testing.T) {
	// check minute range
	_, err := parseTime("00:60:00")
	if err == nil {
		t.Error("Expected error minute not within valid range, got nil")
	}
	// check string as minute
	_, err = parseTime("00:b:00")
	if err == nil {
		t.Error("Expected error minute not a number, got nil")
	}
}

func TestInvalidSecond(t *testing.T) {
	// check second range
	_, err := parseTime("00:00:60")
	if err == nil {
		t.Error("Expected error second not within valid range, got nil")
	}
	// check string as second
	_, err = parseTime("00:00:c")
	if err == nil {
		t.Error("Expected error second not a number, got nil")
	}
}

func TestValidInput(t *testing.T) {
	// check valid input
	time, err := parseTime("12:34:56")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if time.hour != 12 || time.minute != 34 || time.second != 56 {
		t.Errorf("Expected 12:34:56, got %v", time)
	}
}

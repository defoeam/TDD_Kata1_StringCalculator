package string_calculator

import (
	"errors"
	"strconv"
	"strings"
)

type stringCalculator struct {
	calledCount int
}

// newStringCalculator creates a new instance of stringCalculator.
func newStringCalculator() *stringCalculator {
	return &stringCalculator{
		calledCount: 0,
	}
}

// Add takes a string of numbers and returns their sum.
func (sc *stringCalculator) add(numbers string) (int, error) {
	// Base case
	if numbers == "" {
		return 0, nil
	}

	// Parse the numbers from the string
	nums := parseNumStrings(numbers)

	// Interate through the split numbers and convert them to integers
	sum := 0
	issues := make([]string, 0)
	for _, num := range nums {
		num, _ := strconv.Atoi(num)

		// Check for negative numbers
		if num < 0 {
			issues = append(issues, strconv.Itoa(num))
			continue
		}

		// Check for numbers greater than 1000
		if num > 1000 {
			continue
		}

		sum += num
	}

	// If there are any issues, return an error
	if len(issues) > 0 {
		return 0, errors.New("negative numbers are not allowed: " + strings.Join(issues, ", "))
	}

	// Increment the called count
	sc.calledCount++

	// Return the sum
	return sum, nil
}

// parseNumStrings takes a string of numbers and returns a slice of strings.
func parseNumStrings(numbers string) []string {
	// Check for custom delimiters
	if strings.HasPrefix(numbers, "//") {
		// “//[delimiter]\n[numbers…]”
		// Example: “//;\n1;2” == 3
		// Any length delimiter example: “//[***]\n1***2” == 3
		// “//[*][%]\n1*2%3” == 6

		// Multi-character delimiter
		if strings.HasPrefix(numbers, "//[") {
			delimiters := make([]string, 0)

			// Find how many delimiters are present
			delimiterCount := strings.Count(numbers, "[")

			// Find each delimiter
			temp := numbers
			for i := 0; i < delimiterCount; i++ {
				temp = strings.SplitN(temp, "[", 2)[1]
				delimiter := strings.SplitN(temp, "]", 2)[0]
				delimiters = append(delimiters, delimiter)
			}

			// Replace all delimiters with a comma
			temp = strings.SplitN(temp, "\n", 2)[1]
			for _, delimiter := range delimiters {
				temp = strings.ReplaceAll(temp, delimiter, ",")
			}
			return strings.Split(temp, ",")

		} else {
			// Single character delimiter
			delimiter := numbers[2:3]
			numbers = numbers[4:]

			// Replace the delimiter with a comma
			numbers = strings.ReplaceAll(numbers, delimiter, ",")
			return strings.Split(numbers, ",")
		}
	} else {
		// No custom delimiter

		// Replace newlines with commas
		numbers = strings.ReplaceAll(numbers, "\n", ",")

		// Split by commas
		return strings.Split(numbers, ",")
	}
}

// getCalledCount returns the number of times the add method has been called.
func (sc *stringCalculator) getCalledCount() int {
	return sc.calledCount
}

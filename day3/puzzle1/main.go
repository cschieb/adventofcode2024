package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Result should be 161 - 2*4 + 5*5 + 11*8 + 8*5
const testInput = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func main() {
	// Find the valid multiply functions from the input string and add them together to get the answer
	// A valid multiply function is mul(X, Y) where X and Y are both 1-3 digit numbers.
	// Invalid characters should be ignored.

	// Build the regex to parse valid commands
	rx, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	if err != nil {
		panic(err)
	}

	input, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	matches := rx.FindAllString(input, -1)

	fmt.Printf("All matches to the regex are: %v\n", matches)

	sum, err := multiplyAndAdd(matches)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The sum of all valid multiplication functions is %d", sum)
}

// parseInputFile reads the file specified by the provided filename/path as a single string and returns it.
func parseInputFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}

// multiplyAndAdd takes a list of strings, where each string is in the format specified above
// and parses them, performing multiplication actions and returns the sum of the multiplications.
func multiplyAndAdd(matches []string) (int, error) {
	var sum int
	for _, match := range matches {
		indexOpenParen := strings.Index(match, "(")
		indexCloseParen := strings.Index(match, ")")
		indexComma := strings.Index(match, ",")

		firstTerm, err := strconv.Atoi(match[indexOpenParen+1 : indexComma])
		if err != nil {
			return -1, fmt.Errorf("unable to convert first term of %s to integer due to: %w", match, err)
		}

		secondTerm, err := strconv.Atoi(match[indexComma+1 : indexCloseParen])
		if err != nil {
			return -1, fmt.Errorf("unable to convert second term of %s to integer due to: %w", match, err)
		}

		sum += firstTerm * secondTerm
	}

	return sum, nil
}

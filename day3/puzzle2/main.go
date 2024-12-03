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
const testInput = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func main() {
	// Find the valid multiply functions from the input string and add them together to get the answer
	// A valid multiply function is mul(X, Y) where X and Y are both 1-3 digit numbers.
	// Invalid characters should be ignored.
	// Additionally, do() and don't() commands should be parsed and handled. A don't() command should cause
	// all mul() functions after it to be ignored from the total sum until a do() command is encountered.

	// Build the regex to parse valid commands
	rx, err := regexp.Compile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)

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

const doCommand = "do()"
const dontCommand = "don't()"

// multiplyAndAdd takes a list of strings, where each string is in the format specified above
// and parses them, performing multiplication actions and returns the sum of the multiplications.
// Also has special handling for do() and don't() commands, which control enabling and disabling of multiplication operations.
func multiplyAndAdd(matches []string) (int, error) {
	var sum int
	include := true
	for _, match := range matches {
		if match == doCommand {
			include = true
			continue
		} else if match == dontCommand {
			include = false
			continue
		}

		if !include {
			continue
		}
		// After this point, we know we have a mul() command and can safely parse it.

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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Increasing = iota
	Decreasing
)

func main() {
	// Given a list of lists of integers, determine the number of
	// "safe" lists. A list is "safe" on if BOTH of the following are true:
	// 1. The integers in the list are either all increasing or decreasing
	// 2. Adjacent values differ by at least 1 and at most 3
	//
	// Additionally, one "bad" value in a list can be ignored, and the list will still
	// be safe.
	input, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	safeCount := 0
	for _, ints := range input {
		if isSafe(ints) {
			safeCount++
		}
	}

	fmt.Printf("There are %d safe readings in the input file.\n", safeCount)
}

// isSafe tests the "safety" of a list of numbers, as defined above.
func isSafe(elements []int) bool {
	// Each list gets a single error tolerance
	tolerance := 1
	var direction int
	for i, j := 0, 1; i < len(elements)-1 && j < len(elements); {
		// First iteration, determine if we are increasing or decreasing
		if i == 0 {
			if elements[i] > elements[j] {
				direction = Decreasing
			} else if elements[i] < elements[j] {
				direction = Increasing
			} else {
				// Equivalent, not safe
				return false
			}
		}

		var difference int
		if direction == Increasing {
			if elements[i] >= elements[j] {
				return false
			}
			difference = elements[j] - elements[i]
			// Don't need to check for 1, we already know the numbers aren't even
			if difference > 3 {
				return false
			}
		} else if direction == Decreasing {
			if elements[i] <= elements[j] {
				return false
			}
			difference = elements[i] - elements[j]
			// Don't need to check for 1, we already know the numbers aren't even
			if difference > 3 {
				return false
			}
		}

		i++
		j++
	}
	return true
}

// parseInputFile reads in the input file in the below format, a slice of slice of ints:
// 1 2 3 4 5 6
// 11 34 45 65 98
// 43 65 78 9 2
func parseInputFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([][]int, 0)

	// Using bufio to read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Parsed line: %s\n", line)

		// Each line has values that are delimited by a space, so split it
		elements := strings.Split(line, " ")

		convertedElements := make([]int, len(elements))

		for i, element := range elements {
			converted, err := strconv.Atoi(element)
			if err != nil {
				return nil, fmt.Errorf("error parsing element %s due to: %w", element, err)
			}
			convertedElements[i] = converted
		}

		result = append(result, convertedElements)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file due to : %w", err)
	}

	return result, nil
}

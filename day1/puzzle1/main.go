package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Given 2 lists of numbers, we need to find the difference between the elements in the list
	// after they are sorted, and add these differences together to find the total.
	//
	// 1. Read in the lists of numbers from the input file
	// 2. Sort the lists
	// 3. Find the difference between the elements in the list after they are sorted (make sure to take absolute value)
	// 4. Add these differences together to find the total

	firstList, secondList, err := parseInputFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Sort the lists
	sort.Ints(firstList)
	sort.Ints(secondList)

	// Now, simply iterate over the lists and find the difference between the elements (taking the absolute value)
	// and add them together to get the result.
	// Note that we assume the lists are the same length.
	totalDifference := 0

	for i := 0; i < len(firstList); i++ {
		difference := firstList[i] - secondList[i]
		// Doing this instead of using math.Abs to avoid conversion to/from float type
		if difference < 0 {
			difference = -difference
		}
		totalDifference += difference
	}

	fmt.Printf("The total difference between the lists is %d", totalDifference)
}

// parseInputFile reads in the input file in the below format, and returns the 2 lists of numbers:
// 1234   5678
// 4321   8765
// 1234   5678
// 4321   8765
func parseInputFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	// Using bufio to read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Parsed line: %s\n", line)

		// Each line has values that are delimited by 3 spaces, so split it
		elements := strings.Split(line, "   ")

		firstElem, err := strconv.Atoi(elements[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing first element %s due to : %w", elements[0], err)
		}
		firstList = append(firstList, firstElem)

		secondElem, err := strconv.Atoi(elements[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing second element %s due to : %w", elements[1], err)
		}
		secondList = append(secondList, secondElem)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input file due to : %w", err)
	}

	return firstList, secondList, nil
}

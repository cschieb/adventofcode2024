package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Given 2 lists of numbers, we need to calculate a "similarity score" for the 2 lists. This score is defined by
	// multiplying each element in the left list by the number of times it occurs in the right list, then adding all
	// of the results together.
	//
	// 1. Read in the lists of numbers from the input file
	// 2. Iterate over the second list, and create a map of int to int for keeping track of the value -> times seen
	// 3. Iterate over the first list, and calculate the total similarity score by looking up the multiplier from the
	//    map created in step 2

	// Read in the lists
	firstList, secondList, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	// Determine # of times seen for each value in right list
	occurences := make(map[int]int, len(secondList))
	for _, element := range secondList {
		occurences[element] = occurences[element] + 1
	}

	// Iterate over first list and calculate total similarity score
	similarityScore := 0
	for _, element := range firstList {
		similarityScore += element * occurences[element]
	}

	fmt.Printf("The total similarity score for the 2 lists is: %d", similarityScore)
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
			return nil, nil, fmt.Errorf("Error parsing first element %s due to : %w", elements[0], err)
		}
		firstList = append(firstList, firstElem)

		secondElem, err := strconv.Atoi(elements[1])
		if err != nil {
			return nil, nil, fmt.Errorf("Error parsing second element %s due to : %w", elements[1], err)
		}
		secondList = append(secondList, secondElem)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error reading input file due to : %w", err)
	}

	return firstList, secondList, nil
}

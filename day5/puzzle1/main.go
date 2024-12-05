package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Given a set of ordering rules and print orders, calculate the sum of the middle values of all
	// valid orders.
	//
	// 75,47,61,53,29
	// 75 - Can't have seen a 29, 53, 47, 61, 13
	// 47 - Can't have seen a 13, 61, 29
	// 61 - Can't have seen a 13, 53, 29
	// 53 - Can't have seen a 29, 13
	// 29 - Can't have seen a 13
	//
	// We only really care about what we have seen in the past, so we just need to keep track
	// of everything we have seen in a map (for quick lookup) and walk through the values, checking
	// memory and storing as we go
	orderRules, printOrders, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The parsed ordering rules are: %v\n", orderRules)
	fmt.Printf("The parsed printing orders are: %v\n", printOrders)

	sumOfValidMiddles := 0

	for _, order := range printOrders {
		sumOfValidMiddles += validatePrintOrder(order, orderRules)
	}

	fmt.Printf("The sum of the valid print orders is %d", sumOfValidMiddles)
}

// validatePrintOrder validates the provided print order against the supplied rules.
// If the order is valid, it returns the middle page number in the order (otherwise, returns 0).
func validatePrintOrder(printOrder []int, rules map[int][]int) int {
	// Since go doesn't have a set type, using a map to emulate one (i'm too lazy to create a whole type definition right now)
	previous := make(map[int]bool, len(printOrder))

	for _, pageNum := range printOrder {
		rulesForPage, ok := rules[pageNum]
		if ok {
			// For each number in the "rules", make sure that we did not already see it in the past
			for _, rule := range rulesForPage {
				if _, ok := previous[rule]; ok {
					// We saw this in the past which means the rule is violated! Fail
					return 0
				}
			}
		}

		previous[pageNum] = true
	}

	// If we've made it here, it means the order passed all the rules.
	return printOrder[(len(printOrder)-1)/2]
}

// parseInputFile parses the input file line by line, and parsing the ordering rules and printing
// order. The input file will be in the following format:
// 47|53
// 97|13
// 97|61
// 97|47
// 75|29
// 61|13
// 75|53
// 29|13
// 97|29
// 53|29
// 61|53
// 97|53
// 61|29
// 47|13
// 75|47
// 97|75
// 47|61
// 75|61
// 47|29
// 75|13
// 53|13

// 75,47,61,53,29
// 97,61,53,29,13
// 75,29,13
// 75,97,47,61,53
// 61,13,29
// 97,13,75,29,47
//
// The function returns the page ordering rules as a map, and the printing orders as a slice of
// int slices
func parseInputFile(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	orderingRules := make(map[int][]int, 0)
	printOrders := make([][]int, 0)
	allRulesParsed := false

	// Using bufio to read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Parsed line: %s\n", line)

		// This is the empty line separator between the ordering rules and the printing orders
		if strings.TrimSpace(line) == "" {
			allRulesParsed = true
			continue
		}

		// We are in first section - parse pipe-delimited ordering rules
		if !allRulesParsed {
			ordering := strings.Split(line, "|")

			firstTerm, err := strconv.Atoi(ordering[0])
			if err != nil {
				return nil, nil, fmt.Errorf("unable to parse first term %s as int due to: %w", ordering[0], err)
			}

			secondTerm, err := strconv.Atoi(ordering[1])
			if err != nil {
				return nil, nil, fmt.Errorf("unable to parse second term %s as int due to: %w", ordering[1], err)
			}

			if existing, ok := orderingRules[firstTerm]; ok {
				orderingRules[firstTerm] = append(existing, secondTerm)
			} else {
				orderingRules[firstTerm] = []int{secondTerm}
			}
			continue
		}

		// We have parsed all rules, now parse comma-delimited page number orders
		rawPageNumbers := strings.Split(line, ",")

		pageNumbers := make([]int, len(rawPageNumbers))
		for i, rawNum := range rawPageNumbers {
			converted, err := strconv.Atoi(rawNum)
			if err != nil {
				return nil, nil, fmt.Errorf("unable to parse page number %s as int due to: %w", rawNum, err)
			}
			pageNumbers[i] = converted
		}

		printOrders = append(printOrders, pageNumbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading input file due to : %w", err)
	}

	return orderingRules, printOrders, nil
}

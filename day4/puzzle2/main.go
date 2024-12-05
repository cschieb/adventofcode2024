package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type direction int

// Directional constants
const (
	ANY direction = iota
	UP
	DOWN
	LEFT
	RIGHT
	DIAG_UP_RIGHT
	DIAG_DOWN_RIGHT
	DIAG_UP_LEFT
	DIAG_DOWN_LEFT
)

func main() {
	// Given a matrix of characters, see how many instances of "MAS" arranged in an "X" pattern
	// can be found. Can be written forwards or backwards.

	// Read the input file into a matrix view
	matrix, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The parsed input matrix is: %s\n", matrix)

	// Crawl the matrix for x-mas instances
	count := searchForXmas(matrix)

	fmt.Printf("The total number of x-mas occurences is: %d", count)
}

// searchForXmas searches the provided matrix for the center of an X-MAS (an A character),
// then searches the surrounding area to check if it is truly an X-MAS
func searchForXmas(matrix [][]string) int {
	xmasCount := 0
	for i, row := range matrix {
		for j, char := range row {
			// Search for A
			if char == "A" && isXmas(i, j, matrix) {
				xmasCount++
			}
		}
	}

	return xmasCount
}

// isXmas takes a coordinate of an "A" and searches diagonally around it for 2 M/S characters that qualify it
// as an X-MAS instance. Short circuits if on an edge or if one MAS is not found.
func isXmas(row int, col int, matrix [][]string) bool {
	// Short circuit - if we are at any edge of the matrix, we cannot have an x-mas
	if row < 1 || col < 1 || row == len(matrix)-1 || col == len(matrix[row])-1 {
		return false
	}

	oneMas := (search(DIAG_UP_LEFT, "M", row, col, matrix) && search(DIAG_DOWN_RIGHT, "S", row, col, matrix)) ||
		(search(DIAG_UP_LEFT, "S", row, col, matrix) && search(DIAG_DOWN_RIGHT, "M", row, col, matrix))

	if !oneMas {
		return false
	}

	return (search(DIAG_UP_RIGHT, "M", row, col, matrix) && search(DIAG_DOWN_LEFT, "S", row, col, matrix)) ||
		(search(DIAG_UP_RIGHT, "S", row, col, matrix) && search(DIAG_DOWN_LEFT, "M", row, col, matrix))
}

// search searches in the specified direction in the provided matrix from the given coordinate for the
// provided character.
func search(dir direction, char string, row int, col int, matrix [][]string) bool {
	if dir == DIAG_UP_RIGHT {
		if row < 1 || col == len(matrix[row])-1 {
			// Can't go up/right any more, done
			return false
		}
		return matrix[row-1][col+1] == char
	}

	if dir == DIAG_UP_LEFT {
		if row < 1 || col < 1 {
			// Can't go up/left any more, done
			return false
		}
		return matrix[row-1][col-1] == char
	}

	if dir == DIAG_DOWN_RIGHT {
		if row == len(matrix)-1 || col == len(matrix[row])-1 {
			// Can't go down/right any more, done
			return false
		}
		return matrix[row+1][col+1] == char
	}

	if dir == DIAG_DOWN_LEFT {
		if row == len(matrix)-1 || col < 1 {
			// Can't go down/left any more, done
			return false
		}
		return matrix[row+1][col-1] == char
	}

	panic(fmt.Sprintf("Unrecognized direction value received: %d", dir))
}

// parseInputFile parses the input file line by line, splitting each line into
// a slice of characters. Any errors encountered are returned.
func parseInputFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rows := make([][]string, 0)

	// Using bufio to read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Parsed line: %s\n", line)

		// Split each line into individual characters
		characters := strings.Split(line, "")

		rows = append(rows, characters)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input file due to : %w", err)
	}

	return rows, nil
}

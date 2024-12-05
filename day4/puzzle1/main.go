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

var wordToSearch []string = []string{"X", "M", "A", "S"}

func main() {
	// Given a matrix of characters, see how many instances of "XMAS"
	// can be found. Can be horizontal, vertical, diagonal, or backwards.

	// Read the input file into a matrix view
	matrix, err := parseInputFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The parsed input matrix is: %s\n", matrix)

	count := searchForWord(matrix)

	fmt.Printf("The total number of occurences of the word is: %d", count)
}

// searchForWord searches the provided matrix for all occurrences of a word.
func searchForWord(matrix [][]string) int {
	xmasCount := 0
	for i, row := range matrix {
		for j, char := range row {
			// Search for X
			if char == wordToSearch[0] {
				// Add matches to total
				xmasCount += search(ANY, 1, i, j, matrix)
			}
		}
	}

	return xmasCount
}

// search searches recursively for all remaining letters of the string to find in the provided direction.
// the value returned is the total number of matches it found.
func search(dir direction, charIndex int, row int, col int, matrix [][]string) int {
	// This means we've found a match!
	if charIndex == len(wordToSearch) {
		return 1
	}

	if dir == ANY {
		return search(UP, charIndex, row, col, matrix) +
			search(DOWN, charIndex, row, col, matrix) +
			search(LEFT, charIndex, row, col, matrix) +
			search(RIGHT, charIndex, row, col, matrix) +
			search(DIAG_DOWN_LEFT, charIndex, row, col, matrix) +
			search(DIAG_DOWN_RIGHT, charIndex, row, col, matrix) +
			search(DIAG_UP_LEFT, charIndex, row, col, matrix) +
			search(DIAG_UP_RIGHT, charIndex, row, col, matrix)
	}

	if dir == UP {
		if row < 1 {
			// Can't go up any more, done
			return 0
		}
		if matrix[row-1][col] == wordToSearch[charIndex] {
			return search(UP, charIndex+1, row-1, col, matrix)
		}
		return 0
	}

	if dir == DOWN {
		if row == len(matrix)-1 {
			// Can't go down any more
			return 0
		}
		if matrix[row+1][col] == wordToSearch[charIndex] {
			return search(DOWN, charIndex+1, row+1, col, matrix)
		}
		return 0
	}

	if dir == LEFT {
		if col < 1 {
			// Can't go left
			return 0
		}
		if matrix[row][col-1] == wordToSearch[charIndex] {
			return search(LEFT, charIndex+1, row, col-1, matrix)
		}
		return 0
	}

	if dir == RIGHT {
		// ci = 1, row = 0, col = 0,
		if col == len(matrix[row])-1 {
			// Can't go right
			return 0
		}
		if matrix[row][col+1] == wordToSearch[charIndex] {
			return search(RIGHT, charIndex+1, row, col+1, matrix)
		}
		return 0
	}

	if dir == DIAG_UP_RIGHT {
		if row < 1 || col == len(matrix[row])-1 {
			// Can't go up/right any more, done
			return 0
		}
		if matrix[row-1][col+1] == wordToSearch[charIndex] {
			return search(DIAG_UP_RIGHT, charIndex+1, row-1, col+1, matrix)
		}
		return 0
	}

	if dir == DIAG_UP_LEFT {
		if row < 1 || col < 1 {
			// Can't go up/left any more, done
			return 0
		}
		if matrix[row-1][col-1] == wordToSearch[charIndex] {
			return search(DIAG_UP_LEFT, charIndex+1, row-1, col-1, matrix)
		}
		return 0
	}

	if dir == DIAG_DOWN_RIGHT {
		if row == len(matrix)-1 || col == len(matrix[row])-1 {
			// Can't go down/right any more, done
			return 0
		}
		if matrix[row+1][col+1] == wordToSearch[charIndex] {
			return search(DIAG_DOWN_RIGHT, charIndex+1, row+1, col+1, matrix)
		}
		return 0
	}

	if dir == DIAG_DOWN_LEFT {
		if row == len(matrix)-1 || col < 1 {
			// Can't go down/left any more, done
			return 0
		}
		if matrix[row+1][col-1] == wordToSearch[charIndex] {
			return search(DIAG_DOWN_LEFT, charIndex+1, row+1, col-1, matrix)
		}
		return 0
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

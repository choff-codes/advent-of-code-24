package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile() []string {
	file, _ := os.Open("input.txt")
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a slice to store lines
	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line to the slice
		lines = append(lines, scanner.Text())
	}

	return lines
}

// AppendTo2DSlice appends a value to the inner slice at the specified index.
// If the inner slice doesn't exist, it initializes it.
func AppendTo2DSlice(slice [][]int, row int, value int) [][]int {
	// Ensure the outer slice has enough rows
	for len(slice) <= row {
		slice = append(slice, []int{}) // Append an empty slice if needed
	}

	// Append the value to the specified row
	slice[row] = append(slice[row], value)
	return slice
}

// For direction, up = 0, down = 1
// For direction, left = 0, right = 1
// Starting value is the position within the slice (i.e. the other coordinate)
// Index is the column or row index you'd like to iterate through within the database
func getNearestVal(database [][]int, startingValue int, index int, direction int) int {
	sliceToNavigate := database[index]
	fmt.Println(sliceToNavigate)

	previousValue := -1
	nextValue := -1

	for _, value := range sliceToNavigate {
		if value < startingValue {
			previousValue = value
			continue
		}

		if value > startingValue {
			nextValue = value
			break
		}
	}

	if direction == 0 {
		return previousValue
	} else {
		return nextValue
	}
}

func main() {
	fileInput := readFile()

	var columnDatabase = [][]int{}
	var rowDatabase = [][]int{}

	startingPositionRow := -1
	startingPositionColumn := -1

	for rowIndex, row := range fileInput {
		for columnIndex, letter := range row {
			if string(letter) == "#" {
				rowDatabase = AppendTo2DSlice(rowDatabase, rowIndex, columnIndex)
				columnDatabase = AppendTo2DSlice(columnDatabase, columnIndex, rowIndex)
			}

			if string(letter) == "^" {
				startingPositionRow = rowIndex
				startingPositionColumn = columnIndex
			}
		}
	}

	i := 0
	totalCoverage := 0

	for {
		// Flip the direction each time (0 -> 1, 1 -> 0, ..)
		// Flip the database each time (column, row, column, ..)
		// Starting position row and column should come from the return
		// Distance covered should also come from the return
		// If either row / column are -1, you know you have escaped the maze and should call it
		getNearestVal(columnDatabase, startingPositionRow, startingPositionColumn, 0)

		// Safety condition
		if i == 1000 {
			break
		}
	}

	// TODO
	// Starting at the starting position, find the next position until it returns 0
	// Keep track of the total area it covers

	// Test demonstrating upward column traversal
	test := getNearestVal(columnDatabase, 4, 1, 0)

	fmt.Println(test)
}

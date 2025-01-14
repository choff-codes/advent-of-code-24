package main

// importing the packages
import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile() []byte {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return data
}

func TransformInputToArray(input []byte) []string {
	stringifiedInput := string(input)

	return strings.Split(stringifiedInput, "\n")
}

func xmasCheck(xmasGrid []string) int {
	gridFinalIndex := len(xmasGrid) - 1

	totalXmas := 0

	for rowIndex, row := range xmasGrid {
		rowFinalIndex := len(row) - 1

		for letterIndex, letter := range row {
			if string(letter) == "A" {
				if rowIndex >= 1 && rowIndex <= gridFinalIndex-1 &&
					letterIndex >= 1 && letterIndex <= rowFinalIndex-1 {
					totalXmas += checkDiagonalXMas(xmasGrid, rowIndex, letterIndex)
				}
			}
			// if string(letter) == "X" {
			// 	if rowIndex >= 3 {
			// 		totalXmas += checkVertical(xmasGrid, rowIndex, letterIndex, true)
			// 	}
			// 	if rowIndex <= gridFinalIndex-3 {
			// 		totalXmas += checkVertical(xmasGrid, rowIndex, letterIndex, false)
			// 	}
			// 	if letterIndex >= 3 {
			// 		totalXmas += checkHorizontal(xmasGrid, rowIndex, letterIndex, true)
			// 	}
			// 	if letterIndex <= rowFinalIndex-3 {
			// 		totalXmas += checkHorizontal(xmasGrid, rowIndex, letterIndex, false)
			// 	}

			// 	// Diagonal down right
			// 	if rowIndex <= gridFinalIndex-3 && letterIndex <= rowFinalIndex-3 {
			// 		totalXmas += checkDiagonals(xmasGrid, rowIndex, letterIndex, 1)
			// 	}
			// 	// Diagonal down left
			// 	if rowIndex <= gridFinalIndex-3 && letterIndex >= 3 {
			// 		totalXmas += checkDiagonals(xmasGrid, rowIndex, letterIndex, 2)
			// 	}
			// 	// Diagonal up right
			// 	if rowIndex >= 3 && letterIndex <= rowFinalIndex-3 {
			// 		totalXmas += checkDiagonals(xmasGrid, rowIndex, letterIndex, 3)
			// 	}
			// 	// Diagonal up left
			// 	if rowIndex >= 3 && letterIndex >= 3 {
			// 		totalXmas += checkDiagonals(xmasGrid, rowIndex, letterIndex, 4)
			// 	}
			// }
		}
	}

	return totalXmas
}

// direction true = up, direction false = down
// func checkVertical(xmasGrid []string, rowIndex int, letterIndex int, direction bool) int {
// 	if direction {
// 		if string(xmasGrid[rowIndex-1][letterIndex]) == "M" &&
// 			string(xmasGrid[rowIndex-2][letterIndex]) == "A" &&
// 			string(xmasGrid[rowIndex-3][letterIndex]) == "S" {
// 			return 1
// 		}
// 	} else {
// 		if string(xmasGrid[rowIndex+1][letterIndex]) == "M" &&
// 			string(xmasGrid[rowIndex+2][letterIndex]) == "A" &&
// 			string(xmasGrid[rowIndex+3][letterIndex]) == "S" {
// 			return 1
// 		}
// 	}

// 	return 0
// }

// direction true = left, direction false = right
// func checkHorizontal(xmasGrid []string, rowIndex int, letterIndex int, direction bool) int {
// 	if direction {
// 		if string(xmasGrid[rowIndex][letterIndex-1]) == "M" &&
// 			string(xmasGrid[rowIndex][letterIndex-2]) == "A" &&
// 			string(xmasGrid[rowIndex][letterIndex-3]) == "S" {
// 			return 1
// 		}
// 	} else {
// 		if string(xmasGrid[rowIndex][letterIndex+1]) == "M" &&
// 			string(xmasGrid[rowIndex][letterIndex+2]) == "A" &&
// 			string(xmasGrid[rowIndex][letterIndex+3]) == "S" {
// 			return 1
// 		}
// 	}

// 	return 0
// }

// Diagonal down right = 1, Diagonal down left = 2
// Diagonal up right = 3, Diagonal up left = 4
// func checkDiagonals(xmasGrid []string, rowIndex int, letterIndex int, direction int) int {
// 	if direction == 1 {
// 		if string(xmasGrid[rowIndex+1][letterIndex+1]) == "M" &&
// 			string(xmasGrid[rowIndex+2][letterIndex+2]) == "A" &&
// 			string(xmasGrid[rowIndex+3][letterIndex+3]) == "S" {
// 			return 1
// 		}
// 	} else if direction == 2 {
// 		if string(xmasGrid[rowIndex+1][letterIndex-1]) == "M" &&
// 			string(xmasGrid[rowIndex+2][letterIndex-2]) == "A" &&
// 			string(xmasGrid[rowIndex+3][letterIndex-3]) == "S" {
// 			return 1
// 		}
// 	} else if direction == 3 {
// 		if string(xmasGrid[rowIndex-1][letterIndex+1]) == "M" &&
// 			string(xmasGrid[rowIndex-2][letterIndex+2]) == "A" &&
// 			string(xmasGrid[rowIndex-3][letterIndex+3]) == "S" {
// 			return 1
// 		}
// 	} else if direction == 4 {
// 		if string(xmasGrid[rowIndex-1][letterIndex-1]) == "M" &&
// 			string(xmasGrid[rowIndex-2][letterIndex-2]) == "A" &&
// 			string(xmasGrid[rowIndex-3][letterIndex-3]) == "S" {
// 			return 1
// 		}
// 	}

// 	return 0
// }

func checkDiagonalXMas(xmasGrid []string, rowIndex int, letterIndex int) int {
	totalMas := 0

	NEDiagonal := string(xmasGrid[rowIndex+1][letterIndex-1]) + "A" +
		string(xmasGrid[rowIndex-1][letterIndex+1])

	SEDiagonal := string(xmasGrid[rowIndex-1][letterIndex-1]) + "A" +
		string(xmasGrid[rowIndex+1][letterIndex+1])

	NWDiagonal := string(xmasGrid[rowIndex+1][letterIndex+1]) + "A" +
		string(xmasGrid[rowIndex-1][letterIndex-1])

	SWDiagonal := string(xmasGrid[rowIndex-1][letterIndex+1]) + "A" +
		string(xmasGrid[rowIndex+1][letterIndex-1])

	if NEDiagonal == "MAS" {
		totalMas++
	}
	if SEDiagonal == "MAS" {
		totalMas++
	}
	if NWDiagonal == "MAS" {
		totalMas++
	}
	if SWDiagonal == "MAS" {
		totalMas++
	}

	if totalMas == 2 {
		return 1
	}

	return 0
}

func main() {
	fileContents := ReadFile()

	fileArrayForm := TransformInputToArray(fileContents)

	fmt.Println(xmasCheck(fileArrayForm))
}

package main

// importing the packages
import (
	"fmt"
	"log"
	"os"

	// "regexp"
	"slices"
	"strconv"
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

func evaluateReportSafety(report []string) bool {
	lastValue, _ := strconv.Atoi(report[len(report)-1])
	firstValue, _ := strconv.Atoi(report[0])

	alreadyMarked := false

	previousValue := 0

	increasing := firstValue < lastValue

	for subIndex := range report {
		if subIndex == 0 {
			value, err := strconv.Atoi(report[subIndex])

			if err == nil {
				previousValue = value
			}
		} else {
			value, err := strconv.Atoi(report[subIndex])

			if err == nil && !alreadyMarked {
				difference := previousValue - value
				previousValue = value

				if (difference > 0 && increasing) || (difference < 0 && !increasing) ||
					difference > 3 || difference < -3 || difference == 0 {
					alreadyMarked = true
				}

			}
		}
	}

	return alreadyMarked
}

func main() {
	fileContents := ReadFile()

	var reports [][]string
	var unsafeReportNumber int

	fileArrayForm := TransformInputToArray(fileContents)

	for _, element := range fileArrayForm {
		numbers := strings.Split(element, " ")

		reports = append(reports, numbers)
	}

	for _, element := range reports {
		maxIndex := len(element)
		resolved := false

		unsafe := evaluateReportSafety(element)

		if unsafe {
			for i := 0; i < maxIndex; i++ {
				if resolved {
					break
				}
				frontHalf := element[0:i]
				backHalf := element[i+1 : maxIndex]

				fullArray := slices.Concat(frontHalf, backHalf)

				fmt.Print("element", element, "\n")
				fmt.Print("frontHalf", frontHalf, "\n")
				fmt.Print("backHalf", backHalf, "\n")
				fmt.Print("fullArray", fullArray, "\n")

				resolved = !evaluateReportSafety(fullArray)
			}
		} else {
			continue
		}

		if !resolved {
			unsafeReportNumber++
		}
	}

	fmt.Print("Safe Reports:\n")
	fmt.Print(1000 - unsafeReportNumber)
}

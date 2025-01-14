package main

// importing the packages
import (
	"fmt"
	"log"
	"os"

	// "regexp"
	//"slices"
	"strconv"
	"strings"
)

func ReadFile1() []byte {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return data
}

func TransformInputToArray1(input []byte) []string {
	stringifiedInput := string(input)

	return strings.Split(stringifiedInput, "\n")
}

func main1() {
	fileContents := ReadFile()

	var reports [][]string

	unsafeReports := 0

	fileArrayForm := TransformInputToArray1(fileContents)

	for _, element := range fileArrayForm {
		numbers := strings.Split(element, " ")

		reports = append(reports, numbers)
	}

	for _, element := range reports {
		lastValue, err1 := strconv.Atoi(element[len(element)-1])
		firstValue, err2 := strconv.Atoi(element[0])

		alreadyMarked := false

		if err1 == nil && err2 == nil {
			previousValue := 0

			increasing := firstValue < lastValue

			for subIndex := range element {
				if subIndex == 0 {
					value, err := strconv.Atoi(element[subIndex])

					if err == nil {
						previousValue = value
					}
				} else {
					value, err := strconv.Atoi(element[subIndex])

					if err == nil && !alreadyMarked {
						difference := previousValue - value
						previousValue = value

						if (difference > 0 && increasing) || (difference < 0 && !increasing) ||
							difference > 3 || difference < -3 || difference == 0 {
							unsafeReports++
							alreadyMarked = true
						}

					}
				}
			}
		}

		if !alreadyMarked {
			fmt.Print(element)
			fmt.Print("\n")
		}
	}

	fmt.Print("Safe Reports:")
	fmt.Print(1000 - unsafeReports)
}

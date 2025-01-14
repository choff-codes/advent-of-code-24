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

// func RegexForColors(input string) []string {
// 	r, _ := regexp.Compile(`\d*\s*green`)

// 	return r.FindAllString(input, -1)
// }

func main() {
	fileContents := ReadFile()

	var s1 []string
	var s2 []string

	var occurences [100000]int

	fileArrayForm := TransformInputToArray(fileContents)

	for _, element := range fileArrayForm {
		numbers := strings.Split(element, "   ")

		n1, n2 := numbers[0], numbers[1]

		s1 = append(s1, n1)
		s2 = append(s2, n2)
	}

	slices.Sort(s1)
	slices.Sort(s2)

	var totalDiff int = 0

	for _, element := range s2 {
		intElement, err := strconv.Atoi(element)

		if err == nil {
			occurences[intElement] = occurences[intElement] + 1
		}
	}

	for index := range s1 {
		num1, err1 := strconv.Atoi(s1[index])
		//num2, err2 := strconv.Atoi(s2[index])

		if err1 == nil {
			// totalDiff += math.Abs(float64(num1) - float64(num2))
			totalDiff += occurences[num1] * num1
		} else {
			fmt.Println("Error:", err1)
			// fmt.Println("Error:", err2)
		}
	}

	//fmt.Print(s1)
	//fmt.Print(s2)

	fmt.Print(totalDiff)
	// fmt.Println(testInput)
}

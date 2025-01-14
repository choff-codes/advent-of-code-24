package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func main() {
	// File reading efforts
	//fileContents := ReadFile()

	//fileArrayForm := TransformInputToArray(fileContents)

	rulesFilter := regexp.MustCompile(`(\d*)\|(\d*)`)
	listFilter := regexp.MustCompile(`(\d*,\d*)*\n`)

	isolatedRules := rulesFilter.FindAllString(adventData, -1)
	isolatedLists := listFilter.FindAllString(adventData, -1)

	cleanedUpLists := []string{}
	fmt.Println("isolatedLists", isolatedLists)

	middleTotal := 0

	ruleDatabase := [100][100]int{}

	// bullshit new line bloat cleanup because i can't just read this in from a file
	for _, element := range isolatedLists {
		if element == "\n" {
			continue
		} else {
			cleanedUpLists = append(cleanedUpLists, element)
		}
	}

	for _, element := range isolatedRules {
		numberFilter := regexp.MustCompile(`\d*`)

		justNumbers := numberFilter.FindAllString(element, -1)

		ruleBase, _ := strconv.Atoi(justNumbers[0])
		ruleSecondary, _ := strconv.Atoi(justNumbers[1])

		ruleDatabase[ruleBase][ruleSecondary] = 1
	}

	for _, list := range cleanedUpLists {
		splitList := strings.Split(list, ",")
		listFinalIndex := len(splitList) - 1

		badList := false

		for i := listFinalIndex; i >= 0; i-- {
			if !badList {
				for j := 0; j <= i; j++ {
					cleanedFirst := strings.ReplaceAll(splitList[i], " ", "")
					cleanedSecond := strings.ReplaceAll(splitList[j], " ", "")
					cleanedFirst = strings.ReplaceAll(cleanedFirst, "\n", "")
					cleanedSecond = strings.ReplaceAll(cleanedSecond, "\n", "")

					base, _ := strconv.Atoi(cleanedFirst)
					secondary, _ := strconv.Atoi(cleanedSecond)

					if ruleDatabase[base][secondary] == 1 {
						badList = true
						break
					}
				}
			}
		}

		fmt.Println(badList)

		if badList {
			middleValue, _ := strconv.Atoi(splitList[len(splitList)/2])
			middleTotal += middleValue
		}
	}
	fmt.Println(cleanedUpLists)

	fmt.Println(middleTotal)
}

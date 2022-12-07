package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"golang.org/x/exp/utf8string"
)

func getItemTypes(bagContents *utf8string.String) [52]bool {
	var compartmentTypes [52]bool

	for i := 0; i < bagContents.RuneCount(); i++ {
		var item rune = bagContents.At(i)
		typeIndex := getTypeIndex(item)
		compartmentTypes[typeIndex] = true

	}
	return compartmentTypes
}

func getTypeIndex(item rune) int {
	var asciiValue int = int(item)

	// ascii A is 65
	// ascii a is 97

	if unicode.IsUpper(item) {
		return asciiValue - 39
	}
	return asciiValue - 97
}

func getCommonItemPriority(firstElfTypes, secondElfTypes, thirdElfTypes [52]bool) int {
	for i := 0; i < 52; i++ {
		if firstElfTypes[i] && secondElfTypes[i] && thirdElfTypes[i] {
			return i + 1 // plus one because the priorities start from 1 in this excersize
		}
	}

	panic("There should always be a badge item!!")
}

func main() {

	file, err := os.Open("../../input.txt")

	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPriority int = 0

	var isMoreAvalable bool

	for {

		// get First Elf input
		isMoreAvalable = scanner.Scan()
		if !isMoreAvalable {
			break
		}
		firstElfItems := utf8string.NewString(scanner.Text())
		firstElfTypes := getItemTypes(firstElfItems)

		scanner.Scan()
		secondElfItems := utf8string.NewString(scanner.Text())
		secondElfTypes := getItemTypes(secondElfItems)

		scanner.Scan()
		thirdElfItems := utf8string.NewString(scanner.Text())
		thirdElfTypes := getItemTypes(thirdElfItems)

		totalPriority += getCommonItemPriority(firstElfTypes, secondElfTypes, thirdElfTypes)

	}

	fmt.Println("TotalPriority:", totalPriority)
}

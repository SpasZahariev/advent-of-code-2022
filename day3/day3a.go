package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"golang.org/x/exp/utf8string"
)

func calculatePriority(bagContents *utf8string.String) int {
	var leftCompartmentTypes [52]bool
	var rightCompartmentTypes [52]bool

	leftIndex := 0
	rightIndex := bagContents.RuneCount() - 1

	for leftIndex < rightIndex {

		//populate left compartment
		var leftItem rune = bagContents.At(leftIndex)
		leftTypeIndex := getTypeIndex(leftItem)
		leftCompartmentTypes[leftTypeIndex] = true

		//populate right compartment
		var rightItem rune = bagContents.At(rightIndex)
		rightTypeIndex := getTypeIndex(rightItem)
		rightCompartmentTypes[rightTypeIndex] = true

		leftIndex++
		rightIndex--
	}

	//get the collision
	for i := 0; i < 52; i++ {
		if leftCompartmentTypes[i] && rightCompartmentTypes[i] {
			return i + 1 // because priorities start from 1 in this excersize
		}
	}

	log.Fatal("There should always be a collision!!")
	return -1
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

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalPriority int = 0

	for scanner.Scan() {

		utf8String := utf8string.NewString(scanner.Text())

		totalPriority += calculatePriority(utf8String)

	}

	fmt.Println("TotalPriority:", totalPriority)
}

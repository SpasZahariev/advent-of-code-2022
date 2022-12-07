package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasOverlap(elf1Start, elf1End, elf2Start, elf2End int) bool {

	// Check for intersections
	if elf1Start <= elf2Start && elf2Start <= elf1End || elf1Start <= elf2End && elf2End <= elf1End {
		return true
	}

	if elf2Start <= elf1Start && elf1Start <= elf2End || elf2Start <= elf1End && elf1End <= elf2End {
		return true
	}

	// no collision
	return false

}

func main() {

	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fullyContainedCount int = 0

	for scanner.Scan() {

		var elf1Start, elf1End, elf2Start, elf2End int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elf1Start, &elf1End, &elf2Start, &elf2End)

		if hasOverlap(elf1Start, elf1End, elf2Start, elf2End) {
			fullyContainedCount += 1
		}
	}

	fmt.Println("The # of times one elf fully contains the other's work: ", fullyContainedCount)
}

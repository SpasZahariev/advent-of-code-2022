package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func areElvesColliding(inputLine string) bool {

	elveResponsibilities := strings.Split(inputLine, ",")

	return hasOverlap(elveResponsibilities[0], elveResponsibilities[1])
}

func hasOverlap(area1, area2 string) bool {

	elf1Start, elf1End := parseElveArea(area1)
	elf2Start, elf2End := parseElveArea(area2)

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

func parseElveArea(area string) (int, int) {
	numbers := strings.Split(area, "-")

	startArea, _ := strconv.Atoi(numbers[0])
	endArea, _ := strconv.Atoi(numbers[1])
	return startArea, endArea
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

		if areElvesColliding(scanner.Text()) {
			fullyContainedCount += 1
		}
	}

	fmt.Println("The # of times one elf fully contains the other's work: ", fullyContainedCount)
}

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

	return isOneContainedInOther(elveResponsibilities[0], elveResponsibilities[1])
}

func isOneContainedInOther(area1, area2 string) bool {

	elf1Start, elf1End := parseElveArea(area1)
	elf2Start, elf2End := parseElveArea(area2)

	// check if elf one fully contains elf two's work
	if elf1Start <= elf2Start && elf1End >= elf2End {
		return true
	}

	// check if elf two fully contains elf one's work
	if elf2Start <= elf1Start && elf2End >= elf1End {
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var maxSeen [3]int

func main() {
	fmt.Println("Hello Day1")

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close the file when the program finishes
	defer file.Close()

	// read the file line by line with this scanner
	scanner := bufio.NewScanner(file)

	maxSeen[0] = 0
	maxSeen[1] = 0
	maxSeen[2] = 0

	currentElfCalories := 0

	for scanner.Scan() {

		if scanner.Text() == "" {

			indexOfMin, minValue := getMin()

			if currentElfCalories > minValue {
				maxSeen[indexOfMin] = currentElfCalories
			}
			currentElfCalories = 0
		} else {
			// append to the calories
			if calories, err := strconv.Atoi(scanner.Text()); err == nil {
				currentElfCalories += calories
			}
		}

	}

	fmt.Println(maxSeen[0])
	fmt.Println(maxSeen[1])
	fmt.Println(maxSeen[2])

	fmt.Println("the top 3 elves have this much calories: ", maxSeen[0]+maxSeen[1]+maxSeen[2])
}

func getMin() (int, int) {

	minValue := maxSeen[0]
	indexOfMin := 0

	for i, v := range maxSeen {
		if v < minValue {
			indexOfMin = i
			minValue = v
		}
	}

	return indexOfMin, minValue
}

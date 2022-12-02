package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getPointsForMyHandSign(yourSign string, outcome string) int {

	switch outcome {
	case "X": // we have to lose here
		switch yourSign {
		case "A": // I have to give C
			return 3
		case "B": // I have to give A
			return 1
		default: // I have to give B
			return 2
		}
	case "Y": // we have to draw here
		switch yourSign {
		case "A": // I have to give A
			return 1
		case "B": // I have to give B
			return 2
		default: // I have to give C
			return 3
		}
	default: // we have to win here
		switch yourSign {
		case "A": // I have to give B
			return 2
		case "B": // I have to give C
			return 3
		default: // I have to give A
			return 1
		}
	}

}

func getPointsForBattle(outcome string) int {

	switch outcome {
	case "X":
		return 0
	case "Y":
		return 3
	default:
		return 6
	}

}

func main() {

	fmt.Println("Hello Day:2")

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// close the file when the program finishes
	defer file.Close()

	// read the file line by line with this scanner
	scanner := bufio.NewScanner(file)

	totalPoints := 0

	for scanner.Scan() {
		signs := strings.Split(scanner.Text(), " ")
		yourSign := signs[0]
		outcome := signs[1]

		totalPoints += getPointsForMyHandSign(yourSign, outcome) + getPointsForBattle(outcome)
	}

	fmt.Println("totalPoints: ", totalPoints)
}

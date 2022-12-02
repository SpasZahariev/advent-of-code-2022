package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getPointsForMyHandSign(handSign string) int {

	switch handSign {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}

}

func getPointsForBattle(yourSign string, mySign string) int {

	switch mySign {
	case "X":
		switch yourSign {
		case "A":
			return 3
		case "B":
			return 0
		default:
			return 6
		}
	case "Y":
		switch yourSign {
		case "A":
			return 6
		case "B":
			return 3
		default:
			return 0
		}
	default:
		switch yourSign {
		case "A":
			return 0
		case "B":
			return 6
		default:
			return 3
		}
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
		mySign := signs[1]

		totalPoints += getPointsForMyHandSign(mySign) + getPointsForBattle(yourSign, mySign)
	}

	fmt.Println("totalPoints: ", totalPoints)
}

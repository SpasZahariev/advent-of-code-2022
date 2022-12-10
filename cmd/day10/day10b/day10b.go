package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cycle uint8 = 1 // Let's skip the original 0 cycle and start indexing from 1 for cleaner code

	carryOver := 0              // any value that should be applied inside a register Right now!
	var registerValues [223]int //adding a bit more capacity in case of edge cases

	registerValues[0] = 1
	registerValues[1] = 1 // The register X stars with the initial value of 1

	for cycle < 221 {

		scanner.Scan()
		lineParts := strings.Split(scanner.Text(), " ")

		currentX := registerValues[cycle-1] + carryOver

		if len(lineParts) == 1 {
			// this must be a "noop" command
			carryOver = 0
			registerValues[cycle] = currentX
			cycle += 1
		} else {
			// this is an "addx" command
			value, _ := strconv.Atoi(lineParts[1])
			carryOver = value
			registerValues[cycle] = currentX // we apply the value after 2 cycles are finished
			registerValues[cycle+1] = currentX
			cycle += 2
		}

	}

	twenty := 20 * registerValues[20]
	sixty := 60 * registerValues[60]
	hundred := 100 * registerValues[100]
	hundredForty := 140 * registerValues[140]
	hundredEighty := 180 * registerValues[180]
	twoTwenty := 220 * registerValues[220]

	fmt.Println("Strengths:", twenty, sixty, hundred, hundredForty, hundredEighty, twoTwenty)
	answer := twenty + sixty + hundred + hundredForty + hundredEighty + twoTwenty

	fmt.Println("sum of signal strengths ", answer)
}

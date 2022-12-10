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

	var registerValueAtEndOfCycle [241]int //adding a bit more capacity in case of edge cases

	registerValueAtEndOfCycle[0] = 1 // The register X stars with the initial value of 1

	for cycle < 241 {

		scanner.Scan()
		lineParts := strings.Split(scanner.Text(), " ")

		currentX := registerValueAtEndOfCycle[cycle-1]

		if len(lineParts) == 1 {
			// this must be a "noop" command
			registerValueAtEndOfCycle[cycle] = currentX
			cycle += 1
		} else {
			// this is an "addx" command
			value, _ := strconv.Atoi(lineParts[1])
			registerValueAtEndOfCycle[cycle] = currentX // we apply the value after 2 cycles are finished
			registerValueAtEndOfCycle[cycle+1] = currentX + value
			cycle += 2
		}

	}

	// to get the value at the beginnign of a cycle -> decrement by 1
	twenty := 20 * registerValueAtEndOfCycle[19]
	sixty := 60 * registerValueAtEndOfCycle[59]
	hundred := 100 * registerValueAtEndOfCycle[99]
	hundredForty := 140 * registerValueAtEndOfCycle[139]
	hundredEighty := 180 * registerValueAtEndOfCycle[179]
	twoTwenty := 220 * registerValueAtEndOfCycle[219]

	fmt.Println("Strengths:", twenty, sixty, hundred, hundredForty, hundredEighty, twoTwenty)
	answer := twenty + sixty + hundred + hundredForty + hundredEighty + twoTwenty

	fmt.Println("sum of signal strengths ", answer)
}

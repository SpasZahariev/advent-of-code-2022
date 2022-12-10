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

	// Part B

	rowCount := 6
	columnCount := 40

	crtDisplay := make([][]string, rowCount)
	for i := range crtDisplay {
		crtDisplay[i] = make([]string, columnCount)
	}

	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			cycleForThisPixel := (row * 40) + column + 1                     // +1 at the end because on CYCLE one we are drawing PIXEL zero (aka on cycle one we draw row0,col0)
			spritePosition := registerValueAtEndOfCycle[cycleForThisPixel-1] //-1 because we want to know the Value of X in the beginning of this cycle (X shows the position of the middle of the sprite)

			if abs(spritePosition-column) <= 1 {
				crtDisplay[row][column] = "##"
			} else {
				crtDisplay[row][column] = "  "
			}
		}
	}

	fmt.Println()
	for i := range crtDisplay {
		currentLine := crtDisplay[i]
		fmt.Println(strings.Join(currentLine[:], ""))
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

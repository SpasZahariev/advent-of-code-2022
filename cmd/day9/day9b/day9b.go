package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RopeEnd struct {
	row    int
	column int
}

func (o *RopeEnd) moveSingleStep(direction string) {

	numberOfSteps := 1
	switch direction {
	case "U":
		o.row = o.row - numberOfSteps
	case "D":
		o.row = o.row + numberOfSteps
	case "L":
		o.column = o.column - numberOfSteps
	default:
		o.column = o.column + numberOfSteps

	}
}

func (o *RopeEnd) isTouching(other *RopeEnd) bool {

	// if they are adjacent, diagonal, or overlapping -> they are touching
	// see example: (Head is in the center 0,0)
	// -1,-1 -1,0 -1,1
	//  0,-1  0,0  0,1
	//  1,-1  1,0  1,1

	if intAbs(o.row-other.row) <= 1 && intAbs(o.column-other.column) <= 1 {
		return true
	}

	return false
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (tail *RopeEnd) catchupStep(head *RopeEnd) {
	if tail.row != head.row && tail.column != head.column {
		tail.diagonalCatchup(head)
		return
	}
	if tail.row == head.row {
		if head.column > tail.column {
			// move right
			tail.column += 1
		} else {
			// move left
			tail.column -= 1
		}
		return
	}
	// otherwise we must move Row-wise
	if tail.column == head.column {
		if head.row > tail.row {
			// move down
			tail.row += 1
		} else {
			// move up
			tail.row -= 1
		}
		return
	}
}

func (tail *RopeEnd) diagonalCatchup(head *RopeEnd) {

	switch {
	case head.row < tail.row && head.column < tail.column:
		tail.row -= 1
		tail.column -= 1
	case head.row < tail.row && head.column > tail.column:
		tail.row -= 1
		tail.column += 1
	case head.row > tail.row && head.column > tail.column:
		tail.row += 1
		tail.column += 1
	default:
		tail.row += 1
		tail.column -= 1
	}
}

func simulateMovements(headMoves []string) {

	// create sandbox that is 1000 x 1000
	var sandbox [][]bool = make([][]bool, 1000)

	for i := range sandbox {
		sandbox[i] = make([]bool, 1000)
	}

	// 0 index will be my HEAD
	var ropeKnots [10]RopeEnd
	for i := range ropeKnots {
		ropeKnots[i] = initizeRopeKnot()
	}

	sandbox[500][500] = true

	for _, instruction := range headMoves {

		parts := strings.Split(instruction, " ")
		remainingMoves, _ := strconv.Atoi(parts[1])

		for remainingMoves > 0 {
			ropeKnots[0].moveSingleStep(parts[0])

			// move the tail until they are touching
			for tailIndex := 1; tailIndex < len(ropeKnots); tailIndex++ {
				previousKnot := ropeKnots[tailIndex-1]
				currentKnot := &ropeKnots[tailIndex]
				// keep making steps with the 9 tails until they touch the Head
				if !currentKnot.isTouching(&previousKnot) {
					currentKnot.catchupStep(&previousKnot)

					// only take note of the last Tail in the list
					if tailIndex == 9 {
						sandbox[currentKnot.row][currentKnot.column] = true // mark the spot as visited by the tail

					}

				}

			}

			remainingMoves--
		}

	}

	// Count the visited places by the tail
	count := 0
	for i := 0; i < len(sandbox[0]); i++ {
		for j := 0; j < len(sandbox[0]); j++ {
			if sandbox[i][j] {
				count++
			}
		}
	}

	fmt.Println("Unique visited places by the tail:", count)
}

func initizeRopeKnot() RopeEnd {
	return RopeEnd{500, 500}
}

func main() {
	file, err := os.Open("../input.txt")
	// file, err := os.Open("../test-input.txt")
	// file, err := os.Open("../test-input-b.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var headMoves []string
	for scanner.Scan() {
		headMoves = append(headMoves, scanner.Text())

	}

	simulateMovements(headMoves)

}

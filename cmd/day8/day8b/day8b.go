package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseRow(rowText string) []uint8 {
	row := []uint8{}

	for _, ch := range strings.Split(rowText, "") {
		// value, _ := strconv.ParseUint(ch, 0, 8)
		value, _ := strconv.Atoi(ch)
		row = append(row, uint8(value))
	}

	return row
}

func calculateScore(i, j int, inputMatrix [][]uint8) int {

	if isEdge(i, j, len(inputMatrix[0])) {
		return 0
	}

	return calculateLeft(i, j, inputMatrix) *
		calculateRight(i, j, inputMatrix) *
		calculateUp(i, j, inputMatrix) *
		calculateDown(i, j, inputMatrix)

}

func isEdge(i, j, sideSize int) bool {
	return i == 0 ||
		j == 0 ||
		i == sideSize-1 ||
		j == sideSize-1
}

func calculateLeft(rowIndex, columnIndex int, inputMatrix [][]uint8) int {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	seeableTrees := 0
	for j := columnIndex - 1; j >= 0; j-- {
		seeableTrees += 1
		if inputMatrix[rowIndex][j] >= treeHeight {
			break
		}
	}

	return seeableTrees
}

func calculateRight(rowIndex, columnIndex int, inputMatrix [][]uint8) int {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	seeableTrees := 0
	for j := columnIndex + 1; j < len(inputMatrix[0]); j++ {
		seeableTrees += 1
		if inputMatrix[rowIndex][j] >= treeHeight {
			break
		}
	}

	return seeableTrees
}

func calculateUp(rowIndex, columnIndex int, inputMatrix [][]uint8) int {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	seeableTrees := 0
	for i := rowIndex - 1; i >= 0; i-- {
		seeableTrees += 1
		if inputMatrix[i][columnIndex] >= treeHeight {
			break
		}
	}

	return seeableTrees
}

func calculateDown(rowIndex, columnIndex int, inputMatrix [][]uint8) int {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	seeableTrees := 0
	for i := rowIndex + 1; i < len(inputMatrix[0]); i++ {
		seeableTrees += 1
		if inputMatrix[i][columnIndex] >= treeHeight {
			break
		}
	}

	return seeableTrees
}

func main() {
	file, err := os.Open("../input.txt")
	// file, err := os.Open("../test-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// get the size of the first row so we can make our matrix
	scanner.Scan()
	var rowText string = scanner.Text()
	sideSize := len(rowText)

	// using uint8 to save on space. The numbers will always be 0-9
	inputMatrix := make([][]uint8, sideSize)

	rowNumbers := parseRow(rowText)

	inputMatrix[0] = append(inputMatrix[0], rowNumbers...)
	// inputMatrix[0] = rowNumbers

	count := 1
	for scanner.Scan() {

		rowText = scanner.Text()
		rowNumbers = parseRow(rowText)

		inputMatrix[count] = append(inputMatrix[count], rowNumbers...)
		count += 1

	}

	// create matrix of visible/invisible trees
	visibleTrees := make([][]bool, sideSize)
	// I need to instantiate the inner slices
	for i := 0; i < sideSize; i++ {
		visibleTrees[i] = make([]bool, sideSize)
	}

	maxScore := 0

	for i := 0; i < sideSize; i++ {

		for j := 0; j < sideSize; j++ {
			currentScore := calculateScore(i, j, inputMatrix)

			if currentScore > maxScore {
				maxScore = currentScore
			}
		}
	}

	fmt.Println("Max Scening score is: ", maxScore)

}

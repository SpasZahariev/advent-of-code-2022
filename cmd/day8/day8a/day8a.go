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

func isEdge(i, j, sideSize int) bool {
	return i == 0 ||
		j == 0 ||
		i == sideSize-1 ||
		j == sideSize-1
}

func isVisible(i, j int, inputMatrix [][]uint8) bool {

	if isEdge(i, j, len(inputMatrix[0])) {
		return true
	}

	return isVisibleFromLeft(i, j, inputMatrix) ||
		isVisibleFromRight(i, j, inputMatrix) ||
		isVisibleFromUp(i, j, inputMatrix) ||
		isVisibleFromDown(i, j, inputMatrix)

}

func isVisibleFromLeft(rowIndex, columnIndex int, inputMatrix [][]uint8) bool {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	for j := 0; j < columnIndex; j++ {
		if inputMatrix[rowIndex][j] >= treeHeight {
			return false
		}
	}

	return true
}

func isVisibleFromRight(rowIndex, columnIndex int, inputMatrix [][]uint8) bool {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	for j := columnIndex + 1; j < len(inputMatrix[0]); j++ {
		if inputMatrix[rowIndex][j] >= treeHeight {
			return false
		}
	}

	return true
}

func isVisibleFromUp(rowIndex, columnIndex int, inputMatrix [][]uint8) bool {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	for i := 0; i < rowIndex; i++ {
		if inputMatrix[i][columnIndex] >= treeHeight {
			return false
		}
	}

	return true
}

func isVisibleFromDown(rowIndex, columnIndex int, inputMatrix [][]uint8) bool {

	treeHeight := inputMatrix[rowIndex][columnIndex]

	for i := rowIndex + 1; i < len(inputMatrix[0]); i++ {
		if inputMatrix[i][columnIndex] >= treeHeight {
			return false
		}
	}

	return true
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

	countVisible, countTotal := 0, 0

	for i := 0; i < sideSize; i++ {

		for j := 0; j < sideSize; j++ {

			if isVisible(i, j, inputMatrix) {
				visibleTrees[i][j] = true
				countVisible += 1
			}
			countTotal += 1
		}
	}

	fmt.Println("Visible trees are: ", countVisible)
	fmt.Println("All trees are: ", countTotal)

}

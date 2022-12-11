package main

import (
	"advent-of-code/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items           *pkg.Queue[int]
	operator        string
	worryFactor     string
	monkeyDivisor   int
	truePath        int
	falsePath       int
	inspectionCount int
}

func parseMonkey(scanner *bufio.Scanner) Monkey {
	items := parseItems(scanner)
	operation, worryFactor := parseWorryOperation(scanner)
	monkeyNum := parseIntAtEnd(scanner)
	truePath, falsePath := parseThrowDirections(scanner)

	return Monkey{
		&items,
		operation,
		worryFactor,
		monkeyNum,
		truePath,
		falsePath,
		0,
	}

}

func parseItems(scanner *bufio.Scanner) pkg.Queue[int] {
	scanner.Scan()
	items := pkg.NewQueue[int]()
	parts := strings.Split(scanner.Text(), ": ")
	textItems := strings.Split(parts[1], ", ")

	for _, v := range textItems {
		item, _ := strconv.Atoi(v)

		items.Enqueue(item)

	}
	return items
}

func parseWorryOperation(scanner *bufio.Scanner) (string, string) {
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	return parts[6], parts[7]
}

func parseIntAtEnd(scanner *bufio.Scanner) int {
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	lastElement := parts[len(parts)-1]
	value, _ := strconv.Atoi(lastElement)
	return value
}

func parseThrowDirections(scanner *bufio.Scanner) (int, int) {
	truePath := parseIntAtEnd(scanner)
	falsePath := parseIntAtEnd(scanner)
	return truePath, falsePath
}

func throwItems(current *Monkey, monkeys *[]Monkey) {

	for !current.items.IsEmpty() {
		item := current.items.Dequeue()
		throwItem(current, item, monkeys)
	}
}

func throwItem(current *Monkey, item int, monkeys *[]Monkey) {
	newWorry := inspect(current, item)

	if newWorry%current.monkeyDivisor == 0 {
		(*monkeys)[current.truePath].items.Enqueue(newWorry)

	} else {
		(*monkeys)[current.falsePath].items.Enqueue(newWorry)
	}

}

func inspect(monkey *Monkey, item int) int {
	monkey.inspectionCount += 1

	var worry int
	if monkey.worryFactor == "old" {
		worry = item
	} else {
		constant, _ := strconv.Atoi(monkey.worryFactor)
		worry = constant
	}

	if monkey.operator == "*" {
		return (item * worry) / 3
	} else {
		return (item + worry) / 3
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	var monkeys []Monkey

	for sc.Scan() {

		if strings.Contains(sc.Text(), "Monkey") {

			monkey := parseMonkey(sc)
			monkeys = append(monkeys, monkey)

		}

	}

	for i := 0; i < 20; i++ {
		for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
			throwItems(&monkeys[monkeyIndex], &monkeys)
		}
	}
	biggest, secondBiggest := 0, 0
	for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
		if monkeys[monkeyIndex].inspectionCount > secondBiggest {
			secondBiggest = monkeys[monkeyIndex].inspectionCount
		}
		if secondBiggest > biggest {
			biggest, secondBiggest = secondBiggest, biggest
		}
	}

	fmt.Println("Monkey business is: ", biggest*secondBiggest)
}

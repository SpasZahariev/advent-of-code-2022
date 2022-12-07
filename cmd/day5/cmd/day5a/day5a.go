package main

import (
	"bufio"
	"day5/pkg"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func setupInitialState(stacks *[9]pkg.Stack) {
	stacks[0].PushAll([]string{"G", "D", "V", "Z", "J", "S", "B"})
	stacks[1].PushAll([]string{"Z", "S", "M", "G", "V", "P"})
	stacks[2].PushAll([]string{"C", "L", "B", "S", "W", "T", "Q", "F"})
	stacks[3].PushAll([]string{"H", "J", "G", "W", "M", "R", "V", "Q"})
	stacks[4].PushAll([]string{"C", "L", "S", "N", "F", "M", "D"})
	stacks[5].PushAll([]string{"R", "G", "C", "D"})
	stacks[6].PushAll([]string{"H", "G", "T", "R", "J", "D", "S", "Q"})
	stacks[7].PushAll([]string{"P", "F", "V"})
	stacks[8].PushAll([]string{"D", "R", "S", "T", "J"})
}

func parseInput(inputLine string) (int, int, int) {
	inputs := strings.Split(inputLine, " ")

	numberOfItems, _ := strconv.Atoi(inputs[1])
	from, _ := strconv.Atoi(inputs[3])
	to, _ := strconv.Atoi(inputs[5])

	return numberOfItems, from, to
}

func main() {
	file, err := os.Open("../../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("===========")
	fmt.Println("Before:")
	fmt.Println("===========")
	fmt.Println("")

	for i := 0; i < 10; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}

	var stacks [9]pkg.Stack
	setupInitialState(&stacks)

	for scanner.Scan() {
		lineText := scanner.Text()
		numberOfItems, from, to := parseInput(lineText)

		for i := 0; i < numberOfItems; i++ {
			fromStack := &stacks[from-1]
			receiverStack := &stacks[to-1]
			element, _ := fromStack.Pop()
			receiverStack.Push(element)
		}
	}

	fmt.Println("")
	fmt.Println("===========")
	fmt.Println("After:")
	fmt.Println("===========")
	fmt.Println("")

	// Finally print all the array Values
	for _, stack := range stacks {
		fmt.Println(strings.Join(stack, "-"))
		// fmt.Printf("%v", stack)
	}
}

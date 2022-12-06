package main

import (
	"bufio"
	"day6/pkg"
	"fmt"
	"log"
	"os"
)

func areAllDifferent(q *pkg.Queue) bool {
	// a, b, c, d := (*q)[0], (*q)[1], (*q)[2], (*q)[3]

	mySet := make(map[rune]struct{})

	var member struct{}

	for _, v := range *q {
		mySet[v] = member
	}

	return len(mySet) == 4

}

func main() {

	file, err := os.Open("../../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var input string = scanner.Text()
	var queue pkg.Queue

	var counter int = 0

	for _, v := range input {

		if len(queue) < 4 {
			queue.Enqueue(v)
		} else {

			if areAllDifferent(&queue) {
				for _, v := range queue {
					fmt.Println(string(v))
				}
				fmt.Println("position is ", counter)
				break
			}

			queue.Dequeue()
			queue.Enqueue(v)

		}

		counter += 1
	}

}

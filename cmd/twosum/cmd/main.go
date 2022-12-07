package main

import "fmt"

type void struct{} // use void as a syntax suger for empty structure
var member void    // a variable we can reuse for the void type

func main() {

	// asdfasdfasdf asdfasdfa asdfasd

	target := 27
	fmt.Println("We are looking for X and Y that will make: ", target)
	input := []int{12, 25, 15, 14, 45}

	set := make(map[int]void)

	for _, value := range input {
		set[target-value] = member
	}

	var isAnswered bool = false
	for _, item := range input {
		_, isAnswered = set[item]
		if isAnswered {
			fmt.Println("We have a winner!: ", item, " + ", target-item)
			break
		}

	}

	// just in case we didn't find anything
	if !isAnswered {
		fmt.Println("We could not make a two sum to hit the target!")
	}
}

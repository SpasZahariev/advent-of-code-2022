package main

import (
	"bufio"
	"day7/pkg"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func createNewNode(currentNode *pkg.Node, lineText string) {
	// we have to add a child
	splitWords := strings.Split(lineText, " ")
	nodeName := splitWords[1]
	nodeType := splitWords[0]

	nodeValue := 0

	if nodeType != "dir" {

		nodeValue, _ = strconv.Atoi(splitWords[0])
	}

	newNode := pkg.Node{
		Key:      nodeName,
		Value:    nodeValue,
		Parent:   currentNode,
		Children: make(map[string]*pkg.Node),
	}
	// fmt.Println("Adding ", nodeName, " to ", currentNode.Key)
	currentNode.AddChild(&newNode)
}

func calculateSize(node *pkg.Node) int {

	if len(node.Children) == 0 {
		return node.Value
	}

	footPrint := 0
	for _, v := range node.Children {
		footPrint += calculateSize(v)

	}
	node.Value = footPrint
	return node.Value
}

func computeLessOrEqualTo100K(node *pkg.Node, sizes *[]int) {

	for _, v := range node.Children {
		computeLessOrEqualTo100K(v, sizes)
	}
	if node.Value <= 100000 && len(node.Children) != 0 {
		fmt.Println(node.Key, " -> ", node.Value)
		*sizes = append((*sizes), node.Value)
	}
}

func main() {
	file, err := os.Open("../../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	sc.Scan() // skip the first cd / line

	rootNode := pkg.Node{
		Key:      "/",
		Value:    0,
		Parent:   nil,
		Children: make(map[string]*pkg.Node),
	}

	currentNode := &rootNode

	for sc.Scan() {
		lineText := sc.Text()

		if strings.Contains(lineText, "$ ls") {
			continue
		}

		switch {
		case lineText == "$ ls":
			continue
		case lineText == "$ cd ..":
			currentNode = currentNode.Parent
		case strings.Contains(lineText, "$ cd "):
			splitWords := strings.Split(lineText, " ")
			dirName := splitWords[2]
			currentNode = currentNode.Children[dirName]
		default:
			createNewNode(currentNode, lineText)
		}

	}

	fmt.Println("Done with scanning! Time to update the sizes")

	calculateSize(&rootNode)

	fileSizes := []int{}

	computeLessOrEqualTo100K(&rootNode, &fileSizes)

	sum := 0
	for _, v := range fileSizes {

		sum += v
	}
	fmt.Println("The sum of FileSizes is:", sum)
}

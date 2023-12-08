package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alext/aoc/helpers"
)

type Node struct {
	Label string
	Left  *Node
	Right *Node
}

func buildGraph(inCh <-chan string) *Node {
	var startNode *Node

	nodes := make(map[string]*Node)
	getNode := func(label string) *Node {
		node, found := nodes[label]
		if !found {
			node = &Node{Label: label}
			nodes[label] = node
		}
		return node
	}
	for line := range inCh {
		var label, left, right string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &label, &left, &right)
		if err != nil {
			log.Fatal(err)
		}
		node := getNode(label)
		node.Left = getNode(left)
		node.Right = getNode(right)
		if node.Label == "AAA" {
			startNode = node
		}
	}

	return startNode
}
func main() {
	inCh := helpers.StreamLines(os.Stdin)

	instructions := <-inCh
	<-inCh // Discard blank line

	start := buildGraph(inCh)

	current := start
	moveCount := 0
outer:
	for {
		for _, i := range instructions {
			switch i {
			case 'L':
				current = current.Left
			case 'R':
				current = current.Right
			default:
				log.Fatalln("Unexpected instruction:", i)
			}
			moveCount++
			if current.Label == "ZZZ" {
				break outer
			}
		}
	}
	fmt.Println("Total moves:", moveCount)
}

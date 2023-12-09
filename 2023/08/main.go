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

func (n Node) String() string {
	return fmt.Sprintf("%s (%s, %s)", n.Label, n.Left.Label, n.Right.Label)
}

func buildGraph(inCh <-chan string) map[string]*Node {
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
	}

	return nodes
}

func moveCount(graph map[string]*Node, instructions string) int {
	current, ok := graph["AAA"]
	if !ok {
		return -1
	}
	moveCount := 0
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
				return moveCount
			}
		}
	}
}

func moveCountWithGhosts(graph map[string]*Node, instructions string) int {
	var currentSet []*Node
	for label, node := range graph {
		if label[2] == 'A' {
			currentSet = append(currentSet, node)
		}
	}
	fmt.Println(currentSet)
	if len(currentSet) == 0 {
		return -1
	}

	moveCount := 0
	for {
		for _, inst := range instructions {
			allZ := true
			moveCount++
			for i := range currentSet {
				switch inst {
				case 'L':
					currentSet[i] = currentSet[i].Left
				case 'R':
					currentSet[i] = currentSet[i].Right
				default:
					log.Fatalln("Unexpected instruction:", inst)
				}
				if currentSet[i].Label[2] != 'Z' {
					allZ = false
				}
			}
			if allZ {
				return moveCount
			}
		}
	}
}

func main() {
	inCh := helpers.StreamLines(os.Stdin)

	instructions := <-inCh
	<-inCh // Discard blank line

	graph := buildGraph(inCh)

	fmt.Println("Total moves:", moveCount(graph, instructions))
	fmt.Println("Total moves with ghosts:", moveCountWithGhosts(graph, instructions))
}

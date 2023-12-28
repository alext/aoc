package main

import (
	"fmt"
	"log"
	"os"
	"slices"

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

type Graph map[string]*Node

func buildGraph(inCh <-chan string) Graph {
	nodes := make(Graph)
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

func (g Graph) countMoves(instructions string, startNode *Node, anyZ bool) (*Node, int) {
	current := startNode
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
			if anyZ && current.Label[2] == 'Z' {
				return current, moveCount
			}
			if current.Label == "ZZZ" {
				return current, moveCount
			}
		}
	}
}

func (g Graph) moveCount(instructions string) int {
	current, ok := g["AAA"]
	if !ok {
		return -1
	}
	_, count := g.countMoves(instructions, current, false)
	return count
}

func (g Graph) nextNCounts(instructions string, startNode *Node, n int) []int {
	var counts = make([]int, n)
	current := startNode
	for i := 0; i < n; i++ {
		current, counts[i] = g.countMoves(instructions, current, true)
	}
	return counts
}

func (g Graph) moveCountWithGhosts(instructions string) int {
	var startingSet []*Node
	for label, node := range g {
		if label[2] == 'A' {
			startingSet = append(startingSet, node)
		}
	}
	fmt.Println(startingSet)
	if len(startingSet) == 0 {
		return -1
	}
	if len(startingSet) == 1 {
		_, count := g.countMoves(instructions, startingSet[0], true)
		return count
	}
	loopLengths := make([]int, 0, len(startingSet))
	for _, node := range startingSet {
		counts := g.nextNCounts(instructions, node, 10)
		wkg := slices.Clone(counts)
		slices.Sort(wkg)
		wkg = slices.Compact(wkg)
		if len(wkg) != 1 {
			fmt.Printf("%s, has non-repeating counts: %v\n", node, counts)
			return -1
		}
		loopLengths = append(loopLengths, wkg[0])
	}
	fmt.Println("Loop lengths:", loopLengths)

	return helpers.LeastCommonMultiple(loopLengths...)
}

func main() {
	inCh := helpers.StreamLines(os.Stdin)

	instructions := <-inCh
	<-inCh // Discard blank line

	graph := buildGraph(inCh)

	fmt.Println("Total moves:", graph.moveCount(instructions))
	fmt.Println("Total moves with ghosts:", graph.moveCountWithGhosts(instructions))
}

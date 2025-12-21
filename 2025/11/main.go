package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Node struct {
	Label       string
	Connections []*Node
}

type Network map[string]*Node

func (n Network) GetOrCreateNode(label string) *Node {
	node, ok := n[label]
	if !ok {
		node = &Node{Label: label}
		n[label] = node
	}
	return node
}

func (n Network) CountRoutesFrom(node *Node) int {
	count := 0
	for _, conn := range node.Connections {
		if conn.Label == "out" {
			count++
			continue
		}
		count += n.CountRoutesFrom(conn)
	}
	return count
}

func main() {
	network := make(Network)

	helpers.ScanLines(os.Stdin, func(line string) {
		label, connections, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalln("Invalid line:", line)
		}
		n := network.GetOrCreateNode(label)
		for _, conn := range strings.Split(connections, " ") {
			nn := network.GetOrCreateNode(conn)
			n.Connections = append(n.Connections, nn)
		}
	})

	fmt.Println("Routes", network.CountRoutesFrom(network["you"]))
}

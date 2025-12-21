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

func (n Network) countRoutes(from *Node, toLabel string, cache map[string]int) int {
	if from == nil {
		return 0
	}
	if count, ok := cache[from.Label]; ok {
		return count
	}
	count := 0
	for _, conn := range from.Connections {
		if conn.Label == toLabel {
			count++
			continue
		}
		count += n.countRoutes(conn, toLabel, cache)
	}
	cache[from.Label] = count
	return count
}

func (n Network) CountRoutes(fromLabel, toLabel string) int {
	from := n[fromLabel]
	cache := make(map[string]int)
	return n.countRoutes(from, toLabel, cache)
}

func (n Network) CountRoutes1() int {
	return n.CountRoutes("you", "out")
}

func (n Network) CountRoutes2() int {
	svr2dac := n.CountRoutes("svr", "dac")
	fmt.Println("svr2dac:", svr2dac)
	svr2fft := n.CountRoutes("svr", "fft")
	fmt.Println("svr2dac:", svr2fft)

	dac2fft := n.CountRoutes("dac", "fft")
	fmt.Println("dac2fft:", dac2fft)
	fft2dac := n.CountRoutes("fft", "dac")
	fmt.Println("fft2dac:", fft2dac)

	dac2out := n.CountRoutes("dac", "out")
	fmt.Println("dac2out:", dac2out)
	fft2out := n.CountRoutes("fft", "out")
	fmt.Println("fft2out:", fft2out)

	count := 0
	count += svr2dac * dac2fft * fft2out
	count += svr2fft * fft2dac * dac2out
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

	fmt.Println("Routes", network.CountRoutes1())
	fmt.Println("Routes2", network.CountRoutes2())
}

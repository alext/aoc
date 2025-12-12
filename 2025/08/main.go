package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/alext/aoc/helpers"
)

func squared(x int) int { return x * x }

type Box struct {
	X, Y, Z int
}

func (b Box) String() string {
	return fmt.Sprintf("(%d,%d,%d)", b.X, b.Y, b.Z)
}

func (b Box) DistanceTo(other Box) int {
	return squared(other.X-b.X) +
		squared(other.Y-b.Y) +
		squared(other.Z-b.Z)
}

type Pair struct {
	A        Box
	B        Box
	Distance int
}

type Circuit struct {
	Boxes map[Box]bool
	Count int
}

func (c *Circuit) AddBox(b Box) {
	c.Boxes[b] = true
	c.Count++
}

func (c *Circuit) AddBoxesFrom(other *Circuit) {
	maps.Insert(c.Boxes, maps.All(other.Boxes))
	c.Count += other.Count
}

type Network struct {
	circuits     map[*Circuit]bool
	boxToCircuit map[Box]*Circuit
}

func (n *Network) MakeConnection(pair Pair) {
	aCircuit := n.boxToCircuit[pair.A]
	bCircuit := n.boxToCircuit[pair.B]
	if aCircuit != nil && bCircuit != nil {
		if aCircuit == bCircuit {
			// Already in the same circuit
			return
		}

		// Merge circuits
		aCircuit.AddBoxesFrom(bCircuit)
		for box := range bCircuit.Boxes {
			n.boxToCircuit[box] = aCircuit
		}
		delete(n.circuits, bCircuit)

	} else if aCircuit != nil {
		aCircuit.AddBox(pair.B)
		n.boxToCircuit[pair.B] = aCircuit
	} else if bCircuit != nil {
		bCircuit.AddBox(pair.A)
		n.boxToCircuit[pair.A] = bCircuit
	} else {
		c := &Circuit{Boxes: make(map[Box]bool)}
		n.circuits[c] = true
		c.AddBox(pair.A)
		c.AddBox(pair.B)
		n.boxToCircuit[pair.A] = c
		n.boxToCircuit[pair.B] = c
	}
}

func (n *Network) OrderedCircuits() []*Circuit {
	return slices.SortedFunc(maps.Keys(n.circuits), func(a, b *Circuit) int {
		return cmp.Compare(b.Count, a.Count)
	})
}

func main() {
	var boxes []Box
	helpers.ScanLines(os.Stdin, func(line string) {
		parts := strings.Split(line, ",")
		boxes = append(boxes, Box{
			X: helpers.MustAtoi(parts[0]),
			Y: helpers.MustAtoi(parts[1]),
			Z: helpers.MustAtoi(parts[2]),
		})
	})

	var possibleConnections []Pair
	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			a := boxes[i]
			b := boxes[j]
			possibleConnections = append(possibleConnections, Pair{
				A:        a,
				B:        b,
				Distance: a.DistanceTo(b),
			})
		}
	}

	slices.SortFunc(possibleConnections, func(a, b Pair) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	fmt.Println("Possible connections:", len(possibleConnections))

	numConnections := 10
	if len(boxes) > 20 {
		numConnections = 1000
	}

	network := &Network{
		circuits:     make(map[*Circuit]bool),
		boxToCircuit: make(map[Box]*Circuit),
	}
	for i := range numConnections {
		pair := possibleConnections[i]
		network.MakeConnection(pair)
	}

	orderedCircuits := network.OrderedCircuits()

	result := 1
	for i := range 3 {
		result *= orderedCircuits[i].Count
	}
	fmt.Println("Result:", result)

	var pair Pair
	for i := numConnections; i < len(possibleConnections); i++ {
		pair = possibleConnections[i]
		network.MakeConnection(pair)

		if len(network.boxToCircuit) == len(boxes) && len(network.circuits) == 1 {
			break
		}
	}
	fmt.Println("Last connected pair:", pair)
	fmt.Println("Wall distance:", pair.A.X*pair.B.X)
}

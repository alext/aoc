package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/alext/aoc/helpers"
)

type Pos = helpers.Pos

type Map struct {
	Squares map[Pos]string
	Start   Pos
	End     Pos
}

func BuildMap(input [][]string) *Map {
	m := &Map{Squares: make(map[Pos]string)}
	for y, row := range input {
		for x, ch := range row {
			if ch == "#" {
				continue
			}
			p := Pos{X: x, Y: y}
			m.Squares[p] = ch
			if y == 0 {
				m.Start = p
			} else if y == len(input)-1 {
				m.End = p
			}
		}
	}
	return m
}

func (m *Map) FollowPath(firstStep, nodePos Pos) (Pos, int) {
	current, previous := firstStep, nodePos
	stepCount := 0
	for {
		neighbours := slices.DeleteFunc(current.Neighbours(), func(p Pos) bool {
			return p == previous || m.Squares[p] == ""
		})
		if len(neighbours) != 1 {
			// We've reached a node
			return current, stepCount
		}
		previous, current = current, neighbours[0]
		stepCount++
	}
}

func (m *Map) NodeExits(nodePos Pos) []Pos {
	var exits []Pos
	checkPos := func(delta Pos, direction string) {
		p := nodePos.Add(delta)
		if ch := m.Squares[p]; ch == "." || ch == direction {
			exits = append(exits, p)
		}
	}
	checkPos(Pos{Y: -1}, "^")
	checkPos(Pos{Y: 1}, "v")
	checkPos(Pos{X: -1}, "<")
	checkPos(Pos{X: 1}, ">")
	return exits
}

type Path struct {
	Start *Node
	End   *Node
	Steps int
}
type Node struct {
	Location Pos
	OutPaths []*Path
}

type Graph struct {
	Nodes []*Node
	Start *Node
	End   *Node
}

func BuildGraph(m *Map) *Graph {
	type PathEntry struct {
		Start     *Node
		FirstStep Pos
	}
	g := &Graph{
		Start: &Node{Location: m.Start},
		End:   &Node{Location: m.End},
	}
	g.Nodes = []*Node{g.Start, g.End}
	seenNodes := map[Pos]*Node{m.Start: g.Start, m.End: g.End}
	todo := []*PathEntry{{
		Start:     g.Start,
		FirstStep: Pos{X: m.Start.X, Y: m.Start.Y + 1},
	}}
	for len(todo) > 0 {
		p := todo[0]
		todo = todo[1:]
		fmt.Println("Considering path starting", p.FirstStep)

		nodePos, stepCount := m.FollowPath(p.FirstStep, p.Start.Location)

		fmt.Println("Reached node at", nodePos, "steps:", stepCount)

		node, existing := seenNodes[nodePos]
		if !existing {
			fmt.Println("Creating new node")
			node = &Node{Location: nodePos}
			g.Nodes = append(g.Nodes, node)
			seenNodes[nodePos] = node
		}
		path := &Path{
			Start: p.Start,
			End:   node,
			Steps: stepCount,
		}
		path.Start.OutPaths = append(path.Start.OutPaths, path)

		if !existing {
			for _, pos := range m.NodeExits(nodePos) {
				fmt.Println("Adding new exit path:", pos)
				todo = append(todo, &PathEntry{
					Start:     node,
					FirstStep: pos,
				})
			}
		}
	}
	return g
}

func (g *Graph) ToDot() {
	file, err := os.OpenFile("graph.dot", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "digraph {")
	for _, node := range g.Nodes {
		fmt.Fprintf(file, "  \"%s\";\n", node.Location)
	}
	for _, node := range g.Nodes {
		for _, path := range node.OutPaths {
			fmt.Fprintf(file, "  \"%s\" -> \"%s\" [label=\"%d\"];\n", node.Location, path.End.Location, path.Steps)
		}
	}
	fmt.Fprintln(file, "}")
}

type Route struct {
	Current    *Node
	Nodes      []*Node
	Paths      []*Path
	TotalSteps int
}

func (r *Route) Clone() *Route {
	clone := &Route{}
	*clone = *r
	clone.Nodes = slices.Clone(r.Nodes)
	clone.Paths = slices.Clone(r.Paths)
	return clone
}

func (r *Route) NextHops() []*Route {
	var routes []*Route
	for _, path := range r.Current.OutPaths {
		next := r.Clone()
		next.Current = path.End
		next.Nodes = append(next.Nodes, path.End)
		next.Paths = append(next.Paths, path)
		next.TotalSteps += path.Steps + 1 // +1 for the node

		routes = append(routes, next)
	}
	return routes
}

func (g *Graph) LongestPath() *Route {
	routes := []*Route{
		{
			Current: g.Start,
			Nodes:   []*Node{g.Start},
			//TotalSteps: 1,
		},
	}
	var bestRoute *Route
	for len(routes) > 0 {
		route := routes[0]
		routes = routes[1:]

		for _, next := range route.NextHops() {
			if next.Current == g.End {
				if bestRoute == nil || next.TotalSteps > bestRoute.TotalSteps {
					bestRoute = next
				}
				continue
			}
			routes = append(routes, next)
		}
	}
	return bestRoute
}

func main() {
	rawMap := BuildMap(helpers.ScanGrid(os.Stdin, ""))
	fmt.Println(rawMap)
	graph := BuildGraph(rawMap)
	fmt.Println(graph)

	//graph.ToDot()

	r := graph.LongestPath()
	fmt.Println("Longest path length:", r.TotalSteps)
}

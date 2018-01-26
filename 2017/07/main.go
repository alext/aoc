package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

var inputRe = regexp.MustCompile(`([a-z]+)\s+\((\d+)\)(?:\s+->\s+([a-z]+(?:,\s+[a-z]+)*))?`)

func main() {
	tower := NewTower()

	helpers.ScanLines(os.Stdin, func(line string) {
		matches := inputRe.FindStringSubmatch(line)
		p := tower.GetProg(matches[1])
		p.Weight = helpers.MustAtoi(matches[2])
		for _, childName := range helpers.SplitCSV(matches[3]) {
			child := tower.GetProg(childName)
			child.Parent = p
			p.Children = append(p.Children, child)
		}
	})
	err := tower.PopulateRoot()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Root program:", tower.Root.Name)
	unbalanced, correction := tower.FindUnbalanced()
	if unbalanced == nil {
		fmt.Println("all appears balanced")
	} else {
		fmt.Printf("Unbalanced program: %s (%d), correction: %d, correct weight: %d\n", unbalanced.Name, unbalanced.Weight, correction, unbalanced.Weight+correction)
	}
}

type Prog struct {
	Name        string
	Weight      int
	StackWeight int
	Parent      *Prog
	Children    []*Prog
}

func (p *Prog) PopulateStackWeight() {
	p.StackWeight = p.Weight
	for _, child := range p.Children {
		child.PopulateStackWeight()
		p.StackWeight += child.StackWeight
	}
}

func (p *Prog) FindUnbalancedChild() (unbalanced *Prog, correction int) {
	if len(p.Children) == 0 {
		return nil, 0
	}
	childWeights := make(map[int][]*Prog)
	for _, child := range p.Children {
		res, weight := child.FindUnbalancedChild()
		if res != nil {
			return res, weight
		}
		childWeights[child.StackWeight] = append(childWeights[child.StackWeight], child)
	}
	if len(childWeights) > 2 {
		panic("Multiple unbalanced children")
	}

	if len(childWeights) == 1 {
		// all balanced
		return nil, 0
	}

	correctStackWeight := 0
	for weight, children := range childWeights {
		if len(children) == 1 {
			unbalanced = children[0]
		}
		if len(children) > 1 {
			correctStackWeight = weight
		}
	}
	return unbalanced, correctStackWeight - unbalanced.StackWeight
}

type Tower struct {
	Progs map[string]*Prog
	Root  *Prog
}

func NewTower() *Tower {
	return &Tower{Progs: make(map[string]*Prog)}
}

func (t *Tower) GetProg(name string) *Prog {
	p, ok := t.Progs[name]
	if !ok {
		p = &Prog{Name: name}
		t.Progs[name] = p
	}
	return p
}

func (t *Tower) PopulateRoot() error {
	var root *Prog
	for _, p := range t.Progs {
		if p.Parent == nil {
			if root != nil {
				return fmt.Errorf("Found multiple roots, %s and %s", root.Name, p.Name)
			}
			root = p
		}
	}
	t.Root = root
	return nil
}

func (t *Tower) FindUnbalanced() (*Prog, int) {
	t.Root.PopulateStackWeight()
	return t.Root.FindUnbalancedChild()
}

package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Rule struct {
	Property string
	Operator string
	Value    int
	Target   string
}

var (
	instRe = regexp.MustCompile(`([a-z]+){(.+?)}`)
	ruleRe = regexp.MustCompile(`([xmas])([><])(\d+):(A|R|[a-z]+)`)
)

func ParseInstructions(in <-chan string) map[string][]Rule {
	instructions := make(map[string][]Rule)
	for line := range in {
		if line == "" {
			break
		}
		// px{a<2006:qkq,m>2090:A,rfg}
		matches := instRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalln("Failed to parse instruction", line)
		}
		label, rules := matches[1], matches[2]
		for _, ruleStr := range strings.Split(rules, ",") {
			var rule Rule
			if !strings.Contains(ruleStr, ":") {
				rule.Target = ruleStr
			} else {
				ruleMatch := ruleRe.FindStringSubmatch(ruleStr)
				if ruleMatch == nil {
					log.Fatalln("Failed to parse rule", ruleStr, "in line", line)
				}
				rule = Rule{
					Property: ruleMatch[1],
					Operator: ruleMatch[2],
					Value:    helpers.MustAtoi(ruleMatch[3]),
					Target:   ruleMatch[4],
				}
			}
			instructions[label] = append(instructions[label], rule)
		}
	}
	return instructions
}

func (r Rule) MatchPart(part Part) string {
	if r.Operator == "" {
		return r.Target
	}
	var partValue int
	switch r.Property {
	case "x":
		partValue = part.X
	case "m":
		partValue = part.M
	case "a":
		partValue = part.A
	case "s":
		partValue = part.S
	default:
		panic("Unexpected property:" + r.Property)
	}
	if r.Operator == ">" {
		if partValue > r.Value {
			return r.Target
		}
	} else if r.Operator == "<" {
		if partValue < r.Value {
			return r.Target
		}
	}
	return ""
}

func (r Rule) ConstrainPath(path *Path) (*Path, *Path) {
	nextPath := &Path{}
	*nextPath = *path // Shallow copy

	nextPath.CurrentTarget = r.Target
	if r.Operator == "" {
		return nextPath, nil
	}

	remainingPath := &Path{}
	*remainingPath = *path

	nextPath.Constraints = maps.Clone(nextPath.Constraints)
	remainingPath.Constraints = maps.Clone(remainingPath.Constraints)

	constraint := nextPath.Constraints[r.Property]
	remainingConstraint := remainingPath.Constraints[r.Property]
	if r.Operator == ">" {
		constraint.Min = max(constraint.Min, r.Value+1)
		remainingConstraint.Max = min(remainingConstraint.Max, r.Value)
	} else if r.Operator == "<" {
		constraint.Max = min(constraint.Max, r.Value-1)
		remainingConstraint.Min = max(remainingConstraint.Min, r.Value)
	}
	nextPath.Constraints[r.Property] = constraint
	remainingPath.Constraints[r.Property] = remainingConstraint
	return nextPath, remainingPath
}

type Part struct {
	X int
	M int
	A int
	S int
}

func ParseParts(in <-chan string) []Part {
	var parts []Part
	for line := range in {
		// {x=787,m=2655,a=1222,s=2876}
		var p Part
		_, err := fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &p.X, &p.M, &p.A, &p.S)
		if err != nil {
			log.Fatalln("Failed to parse part", line, err)
		}
		parts = append(parts, p)
	}
	return parts
}

type System struct {
	Instructions  map[string][]Rule
	AcceptedParts []Part
	RejectedParts []Part
	AcceptedPaths []*Path
}

func (s *System) ProcessPart(part Part) {
	rules := s.Instructions["in"]
	for {
		target := ""
		for _, rule := range rules {
			target = rule.MatchPart(part)
			if target != "" {
				break
			}
		}
		if target == "A" {
			s.AcceptedParts = append(s.AcceptedParts, part)
			return
		} else if target == "R" {
			s.RejectedParts = append(s.RejectedParts, part)
			return
		}
		rules = s.Instructions[target]
	}
}

func (s *System) ProcessParts(parts []Part) {
	for _, part := range parts {
		s.ProcessPart(part)
	}
}

func (s *System) AcceptedRatingTotal() int {
	total := 0
	for _, p := range s.AcceptedParts {
		total += p.X + p.M + p.A + p.S
	}
	return total
}

type Constraint struct {
	Min int
	Max int
}

type Path struct {
	CurrentTarget string
	Constraints   map[string]Constraint
}

func (p Path) String() string {
	return fmt.Sprintln("Target:", p.CurrentTarget, "Constraints:", p.Constraints["x"], p.Constraints["m"], p.Constraints["a"], p.Constraints["s"])
}

func (p Path) Possible() bool {
	for _, c := range p.Constraints {
		if c.Min > c.Max {
			return false
		}
	}
	return true
}

func (s *System) EvaluatePaths() {
	candidates := []*Path{{
		CurrentTarget: "in",
		Constraints: map[string]Constraint{
			"x": {Min: 1, Max: 4000},
			"m": {Min: 1, Max: 4000},
			"a": {Min: 1, Max: 4000},
			"s": {Min: 1, Max: 4000},
		},
	}}

	for len(candidates) > 0 {
		path := candidates[0]
		candidates = candidates[1:]

		//fmt.Print("Considering:", path)

		for _, rule := range s.Instructions[path.CurrentTarget] {
			if path == nil {
				// Should never happen - means we got a rule without constraints that wasn't the last rule
				break
			}
			nextPath, remainingPath := rule.ConstrainPath(path)
			// Evaluate following rules only against constraints that don't match this rule
			path = remainingPath
			if rule.Target == "R" {
				// Nothing more to evaluate from this rule
				continue
			}
			if !nextPath.Possible() {
				fmt.Print("Discarding impossible nextPath:", nextPath)
				continue
			}
			if nextPath.CurrentTarget == "A" {
				s.AcceptedPaths = append(s.AcceptedPaths, nextPath)
				continue
			}
			candidates = append(candidates, nextPath)
		}
	}
}

func (s *System) DistinctTotalRatings() int {
	total := 0
	for _, path := range s.AcceptedPaths {
		combinations := 1
		for _, c := range path.Constraints {
			combinations *= c.Max - (c.Min - 1)
		}
		total += combinations
	}
	return total
}

func main() {
	inCh := helpers.StreamLines(os.Stdin)

	instructions := ParseInstructions(inCh)
	//fmt.Println(instructions)

	parts := ParseParts(inCh)
	//fmt.Println(parts)

	s := &System{Instructions: instructions}
	s.ProcessParts(parts)
	fmt.Println("Accepted:", s.AcceptedParts)
	fmt.Println("Rejected:", s.RejectedParts)
	fmt.Println("AcceptedRatingTotal:", s.AcceptedRatingTotal())

	s.EvaluatePaths()
	fmt.Println(s.AcceptedPaths)
	fmt.Println("Non-Distinct rating combinations:", s.DistinctTotalRatings())
}

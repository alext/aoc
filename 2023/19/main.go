package main

import (
	"fmt"
	"log"
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

func main() {
	inCh := helpers.StreamLines(os.Stdin)

	instructions := ParseInstructions(inCh)
	fmt.Println(instructions)

	parts := ParseParts(inCh)
	fmt.Println(parts)

	s := &System{Instructions: instructions}
	s.ProcessParts(parts)
	fmt.Println("Accepted:", s.AcceptedParts)
	fmt.Println("Rejected:", s.RejectedParts)
	fmt.Println("AcceptedRatingTotal:", s.AcceptedRatingTotal())
}

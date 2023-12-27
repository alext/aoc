package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/alext/aoc/helpers"
)

type Pulse struct {
	Source string
	Value  bool
	Target string
}

type Module struct {
	Kind    string
	Name    string
	Outputs []string
	State   map[string]bool
}

func NewModule(kind, name string, outputs []string) *Module {
	m := &Module{
		Kind:    kind,
		Name:    name,
		Outputs: outputs,
	}
	if kind == "%" {
		m.State = map[string]bool{"": false}
	} else if kind == "&" {
		m.State = make(map[string]bool)
	}
	return m
}

func (m *Module) String() string {
	return fmt.Sprintf("{%s%s %v -> %v}", m.Kind, m.Name, m.State, m.Outputs)
}

func (m *Module) HandlePulse(pulse Pulse) []Pulse {
	switch m.Kind {
	case "": // broadcaster
		return m.sendPulse(pulse.Value)
	case "%":
		return m.handleFlipFlop(pulse)
	case "&":
		return m.handleConjunction(pulse)
	default:
		panic("Unexpected module type: " + m.Kind)
	}
}

func (m *Module) handleFlipFlop(pulse Pulse) []Pulse {
	if pulse.Value {
		// Ignore high pulse
		return nil
	}
	m.State[""] = !m.State[""]
	return m.sendPulse(m.State[""])
}

func (m *Module) handleConjunction(pulse Pulse) []Pulse {
	m.State[pulse.Source] = pulse.Value
	allHigh := true
	for _, value := range m.State {
		if !value {
			allHigh = false
			break
		}
	}
	return m.sendPulse(!allHigh)
}

func (m *Module) sendPulse(value bool) []Pulse {
	var pulses []Pulse
	for _, output := range m.Outputs {
		pulses = append(pulses, Pulse{
			Value:  value,
			Source: m.Name,
			Target: output,
		})
	}
	return pulses
}

type System struct {
	Modules        map[string]*Module
	PulseCountHigh int
	PulseCountLow  int
}

func NewSystem() *System {
	return &System{
		Modules: make(map[string]*Module),
	}
}

func (s *System) ConnectInputs() {
	for _, m := range s.Modules {
		for _, output := range m.Outputs {
			target := s.Modules[output]
			if target == nil {
				continue
			}
			if target.Kind == "&" {
				target.State[m.Name] = false
			}
		}
	}
}

func (s *System) PushButton() {
	pulseQueue := []Pulse{{
		Source: "button",
		Value:  false,
		Target: "broadcaster",
	}}
	for len(pulseQueue) > 0 {
		pulse := pulseQueue[0]
		pulseQueue = pulseQueue[1:]
		if pulse.Value {
			s.PulseCountHigh++
		} else {
			s.PulseCountLow++
		}
		//fmt.Println("Handling pulse:", pulse)

		module := s.Modules[pulse.Target]
		if module == nil {
			continue
		}
		newPulses := module.HandlePulse(pulse)
		//fmt.Println("newPulses:", newPulses)
		if len(newPulses) > 0 {
			pulseQueue = append(pulseQueue, newPulses...)
		}
	}
}

func (s *System) PulseProduct() int {
	return s.PulseCountHigh * s.PulseCountLow
}

func main() {
	// broadcaster -> a, b, c
	// %a -> b
	// &inv -> a
	lineRe := regexp.MustCompile(`^([%&]?)([a-z]+)\s+->\s+(.*)$`)
	s := NewSystem()
	helpers.ScanLines(os.Stdin, func(line string) {
		matches := lineRe.FindStringSubmatch(line)
		if matches == nil {
			log.Fatalln("Failed to parse line", line)
		}
		name := matches[2]
		s.Modules[name] = NewModule(matches[1], name, helpers.SplitCSV(matches[3]))
	})
	s.ConnectInputs()
	fmt.Println(s)

	for i := 0; i < 1000; i++ {
		s.PushButton()
	}
	fmt.Println(s)
	fmt.Println("PulseProduct:", s.PulseProduct())
}

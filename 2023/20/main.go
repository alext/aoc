package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"

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

func (m *Module) Reset() {
	for k := range m.State {
		m.State[k] = false
	}
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

func (s *System) Reset() {
	s.PulseCountHigh = 0
	s.PulseCountLow = 0
	for _, m := range s.Modules {
		m.Reset()
	}
}

func (s *System) PushButton(logRecipient string) []Pulse {
	pulseQueue := []Pulse{{
		Source: "button",
		Value:  false,
		Target: "broadcaster",
	}}
	var loggedPulses []Pulse
	for len(pulseQueue) > 0 {
		pulse := pulseQueue[0]
		pulseQueue = pulseQueue[1:]
		if pulse.Value {
			s.PulseCountHigh++
		} else {
			s.PulseCountLow++
		}
		if pulse.Target == logRecipient {
			loggedPulses = append(loggedPulses, pulse)
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
	return loggedPulses
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
		s.PushButton("")
	}
	fmt.Println(s)
	fmt.Println("PulseProduct:", s.PulseProduct())

	s.Reset()
	fmt.Println(s)
	var rxSender *Module
	for _, m := range s.Modules {
		if slices.Contains(m.Outputs, "rx") {
			if rxSender != nil {
				log.Fatalln("Found multiple modules outputting to rx", rxSender, m)
			}
			rxSender = m
		}
	}
	if rxSender.Kind != "&" {
		log.Fatalln("Expected rx sender to be a conjunction; got:", rxSender)
	}
	fmt.Println("rx Sender module:", rxSender)
	highPulses := make(map[string]int)
	for count := 1; count < 10000; count++ {
		inputPulses := s.PushButton(rxSender.Name)
		for _, pulse := range inputPulses {
			if pulse.Value {
				fmt.Println("High pulse from source", pulse.Source, "at count", count)
				if preCount, ok := highPulses[pulse.Source]; ok {
					if preCount*2 != count {
						log.Fatalln("Expected double count for", preCount, count)
					}
					continue
				}
				highPulses[pulse.Source] = count
			}
		}
		if len(highPulses) == len(rxSender.State) {
			break
		}
	}
	fmt.Println("HighPulses:", highPulses)
	var counts []int
	for _, count := range highPulses {
		counts = append(counts, count)
	}

	fmt.Println("Counts coincide at:", helpers.LeastCommonMultiple(counts...))
}

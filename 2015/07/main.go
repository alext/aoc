package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/alext/aoc/helpers"
)

type Wire struct {
	name     string
	lock     sync.RWMutex
	valueSet bool
	value    uint16
	input    chan uint16
	outputs  []chan uint16
}

func (w *Wire) Sender() chan<- uint16 {
	ch := make(chan uint16, 1)
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.input != nil {
		panic("Duplicate outputs to wire " + w.name)
	}
	w.input = ch
	go w.readInput()
	return ch
}

func (w *Wire) readInput() {
	value := <-w.input
	w.lock.Lock()
	defer w.lock.Unlock()
	w.value = value
	w.valueSet = true
	for _, ch := range w.outputs {
		ch <- w.value
	}
}

func (w *Wire) AddReceiver() <-chan uint16 {
	ch := make(chan uint16, 1)
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.valueSet {
		ch <- w.value
	} else {
		w.outputs = append(w.outputs, ch)
	}
	return ch
}

var wires = make(map[string]*Wire)

func getWire(name string) *Wire {
	// Generate a constant input wire when given an integer
	num, err := strconv.ParseUint(name, 10, 16)
	if err == nil {
		w := &Wire{name: "constant"}
		w.Sender() <- uint16(num)
		return w
	}

	w, ok := wires[name]
	if !ok {
		w = &Wire{name: name}
		wires[name] = w
	}
	return w
}

func unaryOperator(op string, in, out *Wire) {
	inCh := in.AddReceiver()
	outCh := out.Sender()

	go func() {
		outCh <- calculate(op, <-inCh, 0)
	}()
}

func binaryOperator(op string, in1, in2, out *Wire) {
	in1Ch := in1.AddReceiver()
	in2Ch := in2.AddReceiver()
	outCh := out.Sender()

	go func() {
		outCh <- calculate(op, <-in1Ch, <-in2Ch)
	}()
}

func calculate(op string, in1, in2 uint16) uint16 {
	var output uint16
	switch op {
	case "NOT":
		output = ^in1
	case "AND":
		output = in1 & in2
	case "OR":
		output = in1 | in2
	case "LSHIFT":
		output = in1 << in2
	case "RSHIFT":
		output = in1 >> in2
	case "WIRE":
		output = in1
	default:
		panic("Unrecognised operator: " + op)
	}
	return output
}

func main() {
	helpers.ScanLines(os.Stdin, func(line string) {
		tokens := strings.Split(line, " ")
		switch {
		case tokens[1] == "AND" || tokens[1] == "OR" || strings.Contains(tokens[1], "SHIFT"):
			in1 := getWire(tokens[0])
			in2 := getWire(tokens[2])
			out := getWire(tokens[4])
			binaryOperator(tokens[1], in1, in2, out)
		case tokens[0] == "NOT":
			in := getWire(tokens[1])
			out := getWire(tokens[3])
			unaryOperator(tokens[0], in, out)
		case len(tokens) == 3:
			in := getWire(tokens[0])
			out := getWire(tokens[2])
			unaryOperator("WIRE", in, out)
		default:
			log.Fatal("Unrecognised instruction line: ", line)
		}
	})
	for name, wire := range wires {
		fmt.Printf("Wire %s: %d\n", name, <-wire.AddReceiver())
	}
	outputWire := getWire("a")
	fmt.Println("Result on wire a:", <-outputWire.AddReceiver())
}

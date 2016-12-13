package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alext/aoc/helpers"
)

type Receiver interface {
	SendValue(int)
}

type Bot struct {
	ID        int
	in        chan int
	valueLow  int
	valueHigh int
	outHigh   Receiver
	outLow    Receiver
}

var bots = make(map[int]*Bot)

func GetBot(id int) *Bot {
	if b, ok := bots[id]; ok {
		return b
	}
	b := &Bot{
		ID: id,
		in: make(chan int, 2),
	}
	bots[id] = b
	return b
}

func (b *Bot) AddReceiver(position string, r Receiver) {
	switch position {
	case "low":
		if b.outLow != nil {
			log.Fatalf("Bot %d, duplicate %s receiver", b.ID, position)
		}
		b.outLow = r
	case "high":
		if b.outHigh != nil {
			log.Fatalf("Bot %d, duplicate %s receiver", b.ID, position)
		}
		b.outHigh = r
	default:
		log.Fatalln("Unrecofnised receiver position:", position)
	}
	if b.outLow != nil && b.outHigh != nil {
		go b.readLoop()
	}
}

func (b *Bot) SendValue(value int) {
	select {
	case b.in <- value:
		//nothing
	default:
		log.Fatalf("Bot %d receiving value %d failed. Channel full", b.ID, value)
	}
}

func (b *Bot) readLoop() {
	b.valueLow = <-b.in
	b.valueHigh = <-b.in
	if b.valueLow > b.valueHigh {
		b.valueLow, b.valueHigh = b.valueHigh, b.valueLow
	}
	b.outLow.SendValue(b.valueLow)
	b.outHigh.SendValue(b.valueHigh)
}

type Output struct {
	ID    int
	value int
	in    chan int
}

var outputs = make(map[int]*Output)

func GetOutput(id int) *Output {
	if o, ok := outputs[id]; ok {
		return o
	}
	o := &Output{
		ID: id,
		in: make(chan int, 1),
	}
	outputs[id] = o
	return o
}

func (o *Output) SendValue(value int) {
	select {
	case o.in <- value:
		//nothing
	default:
		log.Fatalf("Output %d receiving value %d failed. Channel full", o.ID, value)
	}
}

func (o *Output) ReadValue() int {
	if o.value == 0 {
		o.value = <-o.in
	}
	return o.value
}

func GetReceiver(kind string, id int) (r Receiver) {
	switch kind {
	case "bot":
		r = GetBot(id)
	case "output":
		r = GetOutput(id)
	default:
		log.Fatalln("Unrecognised receiver type:", kind)
	}
	return r
}

func main() {
	helpers.ScanLines(os.Stdin, func(line string) {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "value":
			value := mustAtoi(tokens[1])
			targetID := mustAtoi(tokens[5])
			GetReceiver(tokens[4], targetID).SendValue(value)
		case "bot":
			bot := GetBot(mustAtoi(tokens[1]))
			r := GetReceiver(tokens[5], mustAtoi(tokens[6]))
			bot.AddReceiver(tokens[3], r)
			r = GetReceiver(tokens[10], mustAtoi(tokens[11]))
			bot.AddReceiver(tokens[8], r)
		default:
			log.Fatalln("Unrecognised instruction:", line)
		}
	})
	for _, output := range outputs {
		fmt.Printf("Output %d, value: %d\n", output.ID, output.ReadValue())
	}
	for _, bot := range bots {
		if bot.valueLow == 17 && bot.valueHigh == 61 {
			fmt.Printf("Bot %d compared 17 and 61\n", bot.ID)
		}
	}
	fmt.Println("Product of outputs 0, 1 and 2:", GetOutput(0).ReadValue()*GetOutput(1).ReadValue()*GetOutput(2).ReadValue())
}

func mustAtoi(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		log.Fatalln("Failed to parse number:", in)
	}
	return value
}

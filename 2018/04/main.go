package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/alext/aoc/helpers"
)

const (
	StartsShift = iota
	FallsAsleep
	WakesUp
)

type Entry struct {
	TS     time.Time
	Guard  int
	Action int
}

//[1518-02-21 00:02] Guard #613 begins shift
//[1518-06-28 00:46] falls asleep
//[1518-08-25 00:57] wakes up
func ParseEntry(input string) Entry {
	var e Entry
	var err error
	tokens := strings.Split(input, ` `)
	if len(tokens) < 4 {
		log.Fatal("Failed to parse entry:", input)
	}
	e.TS, err = time.Parse("[2006-01-02 15:04]", tokens[0]+" "+tokens[1])
	if err != nil {
		log.Fatal(err)
	}
	switch tokens[2] {
	case "Guard":
		e.Action = StartsShift
		e.Guard = helpers.MustAtoi(tokens[3][1:])
	case "falls":
		e.Action = FallsAsleep
	case "wakes":
		e.Action = WakesUp
	default:
		log.Fatal("Failed to recognise action in entry:", input)
	}
	return e
}

type GuardLog struct {
	Guard      int
	Minutes    [60]int
	lastAsleep int
}

func (gl *GuardLog) AddEntry(e Entry) {
	switch e.Action {
	case FallsAsleep:
		gl.lastAsleep = e.TS.Minute()
	case WakesUp:
		for i := gl.lastAsleep; i < e.TS.Minute(); i++ {
			gl.Minutes[i]++
		}
	}
}

func (gl *GuardLog) TotalMinutes() int {
	total := 0
	for _, m := range gl.Minutes {
		total += m
	}
	return total
}

func (gl *GuardLog) SleepiestMinute() int {
	minute := 0
	maxSleep := 0
	for i, m := range gl.Minutes {
		if m > maxSleep {
			minute = i
			maxSleep = m
		}
	}
	return minute
}

type GuardLogs map[int]*GuardLog

func (gls GuardLogs) AddEntry(e Entry) {
	gl, ok := gls[e.Guard]
	if !ok {
		gl = &GuardLog{Guard: e.Guard}
		gls[e.Guard] = gl
	}
	gl.AddEntry(e)
}

func (gls GuardLogs) Sleepiest() *GuardLog {
	var sleepiest *GuardLog
	max := 0
	for _, g := range gls {
		if mins := g.TotalMinutes(); mins > max {
			sleepiest = g
			max = mins
		}
	}
	return sleepiest
}

func main() {
	entries := make([]Entry, 0)
	helpers.ScanLines(os.Stdin, func(line string) {
		entries = append(entries, ParseEntry(line))
	})
	sort.Slice(entries, func(i, j int) bool { return entries[i].TS.Before(entries[j].TS) })
	var guard int
	gls := make(GuardLogs)
	for _, e := range entries {
		if e.Action == StartsShift {
			guard = e.Guard
		} else {
			e.Guard = guard
		}
		gls.AddEntry(e)
	}
	sleepiest := gls.Sleepiest()
	fmt.Println("Sleepiest Guard:", sleepiest.Guard)
	fmt.Println("Sleepiest Minute:", sleepiest.SleepiestMinute())
	fmt.Println("Result:", sleepiest.Guard*sleepiest.SleepiestMinute())
}

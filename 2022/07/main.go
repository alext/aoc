package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alext/aoc/helpers"
)

type DirEntry struct {
	FileSize int
	IsDir    bool
	Parent   *DirEntry
	Children map[string]*DirEntry
}

func NewDirectory(parent *DirEntry) *DirEntry {
	return &DirEntry{
		IsDir:    true,
		Parent:   parent,
		Children: make(map[string]*DirEntry),
	}
}

func NewFile(parent *DirEntry, size int) *DirEntry {
	return &DirEntry{
		Parent:   parent,
		FileSize: size,
	}
}

func (d *DirEntry) Size() int {
	if !d.IsDir {
		return d.FileSize
	}
	total := 0
	for _, child := range d.Children {
		total += child.Size()
	}
	return total
}

type CLI struct {
	Root    *DirEntry
	Current *DirEntry
}

var lsRe = regexp.MustCompile(`^(dir|\d+)\s+(\S+)$`)

func (c *CLI) Execute(command string) {
	lines := strings.Split(command, "\n")
	if lines[0][0:2] != "$ " {
		log.Fatalf("Unexpected command line %s. Full input:\n%s\n", lines[0], command)
	}
	op, arg, _ := strings.Cut(lines[0][2:], " ")
	output := lines[1:]
	switch op {
	case "cd":
		switch arg {
		case "/":
			c.Current = c.Root
		case "..":
			if c.Current.Parent == nil {
				log.Fatalf("Cannot cd up from root dir")
			}
			c.Current = c.Current.Parent
		default:
			dir, ok := c.Current.Children[arg]
			if !ok {
				log.Fatalf("%s not found in cwd", arg)
			}
			if !dir.IsDir {
				log.Fatalf("%s is not a directory", arg)
			}
			c.Current = dir
		}
	case "ls":
		for _, line := range output {
			matches := lsRe.FindStringSubmatch(line)
			if matches == nil {
				log.Fatalf("Failed to parse ls output: %s", line)
			}
			if matches[1] == "dir" {
				c.Current.Children[matches[2]] = NewDirectory(c.Current)
			} else {
				c.Current.Children[matches[2]] = NewFile(c.Current, helpers.MustAtoi(matches[1]))
			}
		}
	default:
		log.Fatalf("Unexpected operator %s. Full input:\n%s\n", op, command)
	}
}

func DirTotalSizes(dir *DirEntry, max int) int {
	total := 0
	if s := dir.Size(); s <= max {
		total += s
	}
	for _, child := range dir.Children {
		if child.IsDir {
			total += DirTotalSizes(child, max)
		}
	}
	return total
}

func SmallestDir(dir *DirEntry, minSize int) int {
	smallest := dir.Size()
	if smallest < minSize {
		return -1
	}
	for _, child := range dir.Children {
		if child.IsDir {
			s := SmallestDir(child, minSize)
			if s >= minSize && s < smallest {
				smallest = s
			}
		}
	}
	return smallest
}

func main() {

	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data)-1; i++ {
			if string(data[i:i+2]) == "\n$" {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			if string(data[len(data)-1]) == "\n" {
				return len(data), data[:len(data)-1], bufio.ErrFinalToken
			} else {
				return len(data), data, bufio.ErrFinalToken
			}
		}
		// request more data
		return 0, nil, nil
	}

	cli := &CLI{
		Root: NewDirectory(nil),
	}
	cli.Current = cli.Root
	helpers.ScanWrapper(os.Stdin, splitFunc, func(command string) {
		cli.Execute(command)
	})

	fmt.Println("Total sizes <= 100k", DirTotalSizes(cli.Root, 100_000))

	const (
		fsSize       = 70000000
		requiredFree = 30000000
	)

	totalSize := cli.Root.Size()
	fmt.Println("Total size:", totalSize)
	freeSpace := fsSize - totalSize
	fmt.Println("Free space:", freeSpace)
	needToDelete := requiredFree - freeSpace
	fmt.Println("Need to delete:", needToDelete)

	smallestDirToDelete := SmallestDir(cli.Root, needToDelete)
	fmt.Println("Smallest to delete:", smallestDirToDelete)
}

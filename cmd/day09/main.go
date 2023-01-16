package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var tupleRegex, _ = regexp.Compile("([UDLR]) (\\d+)")

type vector struct {
	dir string
	len uint
}

type location struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var vectors []vector
	for scanner.Scan() {
		line := scanner.Text()
		if !tupleRegex.MatchString(line) {
			panic(fmt.Sprintf("line didn't match regex: %s", line))
		}
		matches := tupleRegex.FindStringSubmatch(line)
		d := matches[1]
		l, _ := strconv.ParseInt(matches[2], 10, 32)
		vectors = append(vectors, vector{d, uint(l)})
	}
	fmt.Printf("%v\n", vectors)

	vc := tailVisitCount(vectors)
	fmt.Printf("tail visit count %d\n", vc)
}

func tailVisitCount(vectors []vector) uint {
	head := location{0, 0}
	tail := location{0, 0}
	tailVisits := make(map[location]uint)
	for _, v := range vectors {
		for i := 0; uint(i) < v.len; i++ {
			head, tail = step(head, tail, v.dir)
			tailVisits[tail] = tailVisits[tail] + 1
		}
	}
	return uint(len(tailVisits))
}

func step(head location, tail location, direction string) (location, location) {
	switch direction {
	case "U":
		head = location{head.x, head.y + 1}
		if abs(tail.y-head.y) > 1 {
			tail = location{head.x, tail.y + 1}
		}
	case "D":
		head = location{head.x, head.y - 1}
		if abs(tail.y-head.y) > 1 {
			tail = location{head.x, tail.y - 1}
		}
	case "L":
		head = location{head.x - 1, head.y}
		if abs(tail.x-head.x) > 1 {
			tail = location{tail.x - 1, head.y}
		}
	case "R":
		head = location{head.x + 1, head.y}
		if abs(tail.x-head.x) > 1 {
			tail = location{tail.x + 1, head.y}
		}
	default:
		panic(fmt.Sprintf("Unknown direction '%s'", direction))
	}
	return head, tail
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

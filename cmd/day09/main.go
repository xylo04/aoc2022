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
	return 0
}

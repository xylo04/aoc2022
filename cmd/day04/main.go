package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var lineRegex, _ = regexp.Compile("(\\d+)-(\\d+),(\\d+)-(\\d+)")

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aMin, aMax, bMin, bMax := parseLine(scanner.Text())
		aContainsB := rangeContains(aMin, aMax, bMin) && rangeContains(aMin, aMax, bMax)
		bContainsA := rangeContains(bMin, bMax, aMin) && rangeContains(bMin, bMax, aMax)
		if aContainsB || bContainsA {
			acc++
		}
	}
	println("Total number of fully contained assignments: ", acc)
}

// rangeContains returns whether the target value is between min and max (inclusive).
func rangeContains(min uint, max uint, target uint) bool {
	return min <= target && target <= max
}

func parseLine(line string) (uint, uint, uint, uint) {
	matches := lineRegex.FindStringSubmatch(line)
	return parseInt(matches[1]), parseInt(matches[2]), parseInt(matches[3]), parseInt(matches[4])
}

func parseInt(s string) uint {
	i, _ := strconv.ParseInt(s, 10, 32)
	return uint(i)
}

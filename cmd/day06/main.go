package main

import (
	"bufio"
	"log"
	"os"
)

const distinctLen = 14

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	signal := scanner.Text()
	println(findMarker(signal))
}

func findMarker(signal string) int {
	for i := distinctLen; i < len(signal); i++ {
		window := signal[i-distinctLen : i]
		if allDifferent(window) {
			return i
		}
	}
	return -1
}

func allDifferent(window string) bool {
	for i := 0; i < len(window)-1; i++ {
		for j := i + 1; j < len(window); j++ {
			if window[i] == window[j] {
				return false
			}
		}
	}
	return true
}

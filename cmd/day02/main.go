package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const rock = 1
const paper = 2
const scissors = 3
const lose = 0
const draw = 3
const win = 6

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strat := scanner.Text()
		acc += score(strat)
	}
	fmt.Printf("Your final score is %d", acc)
}

func score(strat string) int {
	switch strat {
	case "A X":
		return lose + scissors
	case "B X":
		return lose + rock
	case "C X":
		return lose + paper

	case "A Y":
		return draw + rock
	case "B Y":
		return draw + paper
	case "C Y":
		return draw + scissors

	case "A Z":
		return win + paper
	case "B Z":
		return win + scissors
	case "C Z":
		return win + rock
	default:
		panic("unexpected input")
	}
}
